package services

import (
	"fmt"

	"github.com/santhoshbhatti/bookstore_users-api/controllers/users/utils/errors"
	"github.com/santhoshbhatti/bookstore_users-api/domain/users"
)

/*CreateUser is the service to create new user*/
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	fmt.Println(">>>>>>>inside CreateUser service >>>>>>>" + user.FirstName)
	return &user, nil
}
