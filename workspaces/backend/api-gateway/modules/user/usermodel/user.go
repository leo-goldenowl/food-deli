package usermodel

import (
	"errors"

	"api-gateway/common"

	"github.com/google/uuid"
)

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;type:varchar(255);unique_index"`
	Password        string        `json:"-" gorm:"column:password;type:varchar(255);not null"`
	LastName        string        `json:"lastName" gorm:"column:last_name;type:varchar(255);not null"`
	FirstName       string        `json:"firstName" gorm:"column:first_name;type:varchar(255);not null"`
	Role            string        `json:"role " gorm:"column:role;type:varchar(255);not null"`
	Salt            string        `json:"-" gorm:"column:salt;type:varchar(255);"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:jsonb;"`
}

func (u *User) GetUserId() uuid.UUID {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

func (User) TableName() string {
	return "users"
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"-" gorm:"column:password"`
	LastName        string        `json:"lastName" gorm:"column:last_name"`
	FirstName       string        `json:"firstName" gorm:"column:first_name"`
	Role            string        `json:"-" gorm:"column:role"`
	Salt            string        `json:"-" gorm:"column:salt"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

type UserLogin struct {
	Email    string `json:"email" gorm:"column:email;"`
	Password string `json:"password" gorm:"column:password"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

var (
	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
)
