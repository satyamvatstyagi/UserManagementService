package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jinzhu/gorm"
	"github.com/satyamvatstyagi/UserManagementService/pkg/app/models"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/cerr"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/consts"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/instrumentation"
	"go.elastic.co/apm/module/apmgorm/v2"
	"go.elastic.co/apm/v2"
)

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(database *gorm.DB) models.UserRepository {
	return &userRepository{
		database: database,
	}
}

func (u *userRepository) RegisterUser(ctx context.Context, userName string, password string) (string, error) {
	span, ctx := instrumentation.TraceAPMRequest(ctx, "RegisterUser", consts.SpanTypeQueryExecution)
	defer span.End()
	db := apmgorm.WithContext(ctx, u.database)
	localUTCTime := time.Now()
	user := &models.User{
		UserName:  userName,
		Password:  password,
		CreatedAt: localUTCTime,
		UpdatedAt: localUTCTime,
	}

	if err := db.Create(user).Error; err != nil {
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
	span, ctx := instrumentation.TraceAPMRequest(ctx, "GetUserByUserName", consts.SpanTypeQueryExecution)
	defer span.End()
	db := apmgorm.WithContext(ctx, u.database)
	var user models.User
	if err := db.Where("user_name = ?", userName).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			apm.CaptureError(ctx, fmt.Errorf("db error: %s", err.Error())).Send()
			return nil, cerr.NewCustomErrorWithCodeAndOrigin("User not found", cerr.InvalidRequestErrorCode, err)
		}
		apm.CaptureError(ctx, fmt.Errorf("db error: %s", err.Error())).Send()
		return nil, err
	}
	return &user, nil
}
