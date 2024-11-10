package featureUser

import (
	"crypto/sha512"
	"fmt"
	"url-shorts.com/internal/db"
)

const passwordSalt = "u&21Nal)-1"

type userMethods interface {
	db.Iterable[UserItem]
	db.RequestMethods
	findById(uint) error
	findByEmail(string) error
	createUser(email string, password string) error
}

type User interface {
	GetId() uint
	setPassword(string)
	isPasswordValid(string) bool
}

func newUserRequest() userMethods {
	return &userRequest{
		IterableOrigin: db.IterableOrigin[UserItem]{
			Origin: &[]UserItem{},
		},
		Request: db.Request{
			Db: db.GetDb(),
		},
	}
}

type userRequest struct {
	db.IterableOrigin[UserItem]
	db.Request
}

func (u *userRequest) findByEmail(email string) error {
	return u.Db.Where("user_email = ?", email).Find(u.Origin).Error
}

func (u *userRequest) findById(id uint) error {
	return u.Db.Where("id = ?", id).Find(u.Origin).Error
}

func (u *userRequest) createUser(email string, password string) error {
	user := UserItem{
		db.User{
			UserEmail: email,
			IsBan:     false,
		},
	}

	user.setPassword(password)

	u.Origin = &[]UserItem{user}
	return u.Db.Save(u.Origin).Error
}

type UserItem struct {
	db.User
}

func (u *UserItem) TableName() string {
	return "users"
}

func (u *UserItem) GetId() uint {
	return u.ID
}

func (u *UserItem) setPassword(password string) {
	sha512Data := sha512.New()
	sha512Data.Write([]byte(password + u.UserEmail + passwordSalt))
	u.Password = fmt.Sprintf("%x", sha512Data.Sum(nil))
}

func (u *UserItem) isPasswordValid(password string) bool {
	sha512Data := sha512.New()
	sha512Data.Write([]byte(password + u.UserEmail + passwordSalt))
	checkPassword := fmt.Sprintf("%x", sha512Data.Sum(nil))

	return checkPassword == u.Password
}
