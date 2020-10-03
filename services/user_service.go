package services

import (
	"fmt"

	"github.com/santhoshbhatti/bookstore_users-api/controllers/users/utils/errors"
	"github.com/santhoshbhatti/bookstore_users-api/domain/users"
)

/*CreateUser is the service to create new user*/
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	fmt.Println(">>>>>>>inside CreateUser service >>>>>>>" + user.FirstName)
	//this should be part of the user package as this can be reused in may places
	//Other functions where user is passed may also need to validate the user email
	/*user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return nil, errors.NewBadRequest("invalid Email address")
	}*/
	//now one more refactoring making Validate a method -----receiver function of User struct
	//This more appropriate as we have validataion now as part of the User struct itself
	if validErr := user.Validate(); validErr != nil {
		return nil, validErr
	}

	if saveErr := user.Save(); saveErr != nil {
		return nil, saveErr
	}
	return &user, nil
}

/*GetUserByID get user by id*/
func GetUserByID(user *users.User) *errors.RestErr {

	restErr := user.Get()
	return restErr

}
