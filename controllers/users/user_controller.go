package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/santhoshbhatti/bookstore_users-api/controllers/users/utils/errors"
	"github.com/santhoshbhatti/bookstore_users-api/domain/users"
	"github.com/santhoshbhatti/bookstore_users-api/services"
)

/*CreateUser function*/
func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	/*bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//TODO handle error and return
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
	if err := json.Unmarshal(bytes, &user); err != nil {
		//TODO handle json error and return
		return
	}*/
	//this below code ShouldBindJson is equivallent to the commented code
	if err := c.ShouldBindJSON(&user); err != nil {
		//handle unmarshall errors here
		restErr := errors.NewBadRequest("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	fmt.Println("!!!!!!!!!!unmarshalled user >>>>> ", user)
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//handle error and exit
		c.JSON(saveErr.Status, saveErr)
		return
	}
	fmt.Println("created user >>>>>>>>>>>>>>> ", result)
	//c.String(http.StatusNotImplemented, "Implement me!")
	c.JSON(http.StatusCreated, result)
}

/*GetUser function*/
func GetUser(c *gin.Context) {
	userIDStr := c.Param("userid")
	userID, userErr := strconv.ParseInt(userIDStr, 10, 64)
	if userErr != nil {
		c.JSON(http.StatusBadRequest, errors.NewBadRequest(fmt.Sprintf("invalid userID, %s", userIDStr)))
		return
	}
	user := users.User{ID: userID}
	getErr := services.GetUserByID(&user)

	if getErr != nil {
		c.JSON(http.StatusNotFound, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

/*SearchUser function*/
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
