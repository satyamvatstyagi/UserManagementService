package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"

	"github.com/satyamvatstyagi/UserManagementService/pkg/app/models"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/cerr"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/consts"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/mtnapm"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/logger"
	"go.elastic.co/apm/v2"
)

type userRepository struct {
	database *gorm.DB
	logger   *logger.MtnLogger
}

func NewUserRepository(database *gorm.DB, logger *logger.MtnLogger) *userRepository {
	return &userRepository{database: database, logger: logger}
}

func (u *userRepository) RegisterUser(ctx context.Context, userName string, password string) (string, error) {

	localUTCTime := time.Now()
	user := &models.User{
		UserName:  userName,
		Password:  password,
		CreatedAt: localUTCTime,
		UpdatedAt: localUTCTime,
	}

	//for fetching the database query
	statement := u.database.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Create(user)
	})

	instrument := mtnapm.InitGormAPM(ctx, "postgresql", statement)
	defer instrument.GetSpan().End()

	if err := u.database.Create(user).Error; err != nil {
		// Check if err is of type *pgconn.PgError and error code is 23505, which is the error code for unique_violation
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == consts.UniqueViolation {
			apm.CaptureError(ctx, fmt.Errorf("db error: %s", pgErr.Error())).Send()
			return "", cerr.NewCustomErrorWithCodeAndOrigin("User already exists for this user", cerr.InvalidRequestErrorCode, err)
		}
		apm.CaptureError(ctx, fmt.Errorf("db error: %s", err.Error())).Send()
		return "", err
	}

	return user.UUID.String(), nil
}

func (u *userRepository) GetUserByUserName(ctx context.Context, userName string) (*models.User, error) {
	var user models.User

	//for fetching the database query
	statement := u.database.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Where("user_name = ?", userName).First(&user)
	})

	instrument := mtnapm.InitGormAPM(ctx, "postgresql", statement)
	defer instrument.GetSpan().End()

	if err := u.database.Where("user_name = ?", userName).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			apm.CaptureError(ctx, fmt.Errorf("db error: %s", err.Error())).Send()
			return nil, cerr.NewCustomErrorWithCodeAndOrigin("User not found", cerr.InvalidRequestErrorCode, err)
		}
		apm.CaptureError(ctx, fmt.Errorf("db error: %s", err.Error())).Send()
		return nil, err
	}
	return &user, nil
}
