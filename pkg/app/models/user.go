package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID      uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid();unique"`
	UserName  string    `gorm:"index:idx_user_uuid,unique;not null;"`
	Password  string    `gorm:"size:255;not null;" json:"password"`
	CreatedAt time.Time `gorm:"not null;"`
	UpdatedAt time.Time `gorm:"not null;"`
}

type UserRepository interface {
	RegisterUser(userID string, password string) (string, error)
	GetUserByUserName(userName string) (*User, error)
}
