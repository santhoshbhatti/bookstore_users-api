package users

import (
	"strings"

	"github.com/santhoshbhatti/bookstore_users-api/controllers/users/utils/errors"
)

/*User represents the user domain object*/
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

/*Validate for validating user*/
func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequest("invalid Email address")
	}
	return nil
}

/*Copy from this user to otherUser*/
func (user *User) Copy(otherUser *User) {
	otherUser.ID = user.ID
	otherUser.FirstName = user.FirstName
	otherUser.LastName = user.LastName
	otherUser.Email = user.Email
	otherUser.DateCreated = user.DateCreated
}
