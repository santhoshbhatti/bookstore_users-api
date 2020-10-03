package users

import (
	"fmt"

	"github.com/santhoshbhatti/bookstore_users-api/controllers/users/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

/*Get from db the user with userId*/
func (user *User) Get() *errors.RestErr {
	result := userDB[user.ID]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}
	result.Copy(user)
	return nil
}

/*Save the give n user*/
func (user *User) Save() *errors.RestErr {
	if userDB[user.ID] != nil {
		return errors.NewBadRequest(fmt.Sprintf("user %d already exists", user.ID))
	}
	userDB[user.ID] = user
	return nil
}
