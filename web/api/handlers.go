package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"v0rt3x/perimeter/web/ws"
)

type Notification struct {
	Class   string `json:"class"`
	Message string `json:"message"`
}

func Test(c *gin.Context) {
	n := Notification{}

	if c.Bind(&n) == nil {
		ws.GetChannel("notifications").Submit(gin.H{"class": n.Class, "message": n.Message})

		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "ERROR",
		})
	}
}
