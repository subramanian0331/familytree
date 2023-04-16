package services

import (
	"errors"
	"fmt"

	"github.com/subramanian0331/familytree/models"
	"github.com/subramanian0331/familytree/store"
)

type UserService struct {
	pql store.PostgresDB
}

type IUserService interface {
	Create(*models.User) error
	Login(*models.User) error
	Logout(Email string) error
	Delete(Email string) error
	ValidatePassword(Password string, Username string) error
}

// logs in user
func (u *UserService) Login(user *models.User) error {
	email := user.Email
	_, err := u.pql.GetUser(email)
	if err != nil {
		// User doesn't exist in the db.
		return errors.New("user doesn't exist")
	}
	err = u.pql.Login(email)
	if err != nil {
		fmt.Println("couldn't add user to db")
		return err
	}
	return nil
}

//
