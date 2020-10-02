package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

/*StartApplication application */
func StartApplication() {
	mapUrls()
	router.Run(":8080")
}
