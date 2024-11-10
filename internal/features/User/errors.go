package featureUser

import (
	"errors"
)

var (
	ErrorUserNotFound         = errors.New("user not found")
	ErrorUserWrongCredentials = errors.New("wrong credentials")
	ErrorUserBanned           = errors.New("user banned")
)
