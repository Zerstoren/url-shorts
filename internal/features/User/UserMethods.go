package featureUser

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func GetAuthenticateUser(ctx *fiber.Ctx) User {
	user := ctx.Locals("user")

	if user == nil {
		return nil
	}

	return ctx.Locals("user").(User)
}

func GetUserById(id uint) User {
	user := newUserRequest()
	if err := user.findById(id); err != nil {
		log.Errorf("can't find user by id: %v", err)
		return nil
	}

	findedUser, ok := user.GetFirst()
	if !ok {
		log.Error("can't get user from collection")
		return nil
	}

	return findedUser
}

func AuthenticateUserByCredentials(session *session.Session, email string, password string) *string {
	findUserRequest := newUserRequest()

	errorText := ""

	if err := findUserRequest.findByEmail(email); err != nil {
		errorText = "Wrong Credentials"
		return &errorText
	}

	user, ok := findUserRequest.GetFirst()

	if !ok {
		errorText = "Wrong Credentials"
		return &errorText
	}

	if !user.isPasswordValid(password) {
		errorText = "Wrong Credentials"
		return &errorText
	}

	session.Set("user", user.GetId())
	if err := session.Save(); err != nil {
		log.Error("Error saving session: " + err.Error())
	}

	return nil
}

func RegisterNewUser(session *session.Session, email string, password string) *string {
	errorText := ""

	if session.Get("user") != nil {
		errorText = "User already authorized"
		return &errorText
	}

	findUserRequest := newUserRequest()

	if err := findUserRequest.findByEmail(email); err != nil {
		log.Errorf("can't process findByEmail request %v", err)
		errorText = "Can't search user"
		return &errorText
	}

	findUserRequest.ResetDb()
	if err := findUserRequest.createUser(email, password); err != nil {
		log.Errorf("Error saving user: %v", err)
		errorText = "Can't create user"
		return &errorText
	}

	user, ok := findUserRequest.GetFirst()

	if !ok {
		errorText = "Can't serve user, try login"
		return &errorText
	}

	session.Set("user", user.GetId())
	if err := session.Save(); err != nil {
		log.Error("Error saving session: " + err.Error())
	}

	return nil
}
