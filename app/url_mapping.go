package app

import (
	"github.com/santhoshbhatti/bookstore_users-api/controllers/ping"
	"github.com/santhoshbhatti/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)

	router.GET("/users/:userid", users.GetUser)

	//router.GET("/users/search", controllers.SearchUser)
}
