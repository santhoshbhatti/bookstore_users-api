package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*Ping is is function to return pong*/
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
