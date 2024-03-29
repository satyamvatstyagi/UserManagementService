package repository

import (
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/satyamvatstyagi/UserManagementService/pkg/app/models"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/cerr"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/consts"
	"github.com/satyamvatstyagi/UserManagementService/pkg/common/logger"
	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
	logger   *logger.MtnLogger
}

func NewUserRepository(database *gorm.DB, logger *logger.MtnLogger) *userRepository {
	return &userRepository{database: database, logger: logger}
}

func (u *userRepository) RegisterUser(userName string, password string) (string, error) {
	localUTCTime := time.Now()
	user := &models.User{
		UserName:  userName,
		Password:  password,
		CreatedAt: localUTCTime,
		UpdatedAt: localUTCTime,
	}

	if err := u.database.Create(user).Error; err != nil {
		// Check if err is of type *pgconn.PgError and error code is 23505, which is the error code for unique_violation
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == consts.UniqueViolation {
			return "", cerr.NewCustomErrorWithCodeAndOrigin("User already exists for this user", cerr.InvalidRequestErrorCode, err)
		}
		return "", err
	}

	return user.UUID.String(), nil
}

func (u *userRepository) GetUserByUserName(userName string) (*models.User, error) {
	var user models.User
	if err := u.database.Where("user_name = ?", userName).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, cerr.NewCustomErrorWithCodeAndOrigin("User not found", cerr.InvalidRequestErrorCode, err)
		}
		return nil, err
	}
	return &user, nil
}
