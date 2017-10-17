package ws

import (
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

func StreamChannel(c *gin.Context) {
	channelName := c.Param("channel")

	eventListener := OpenListener(channelName)
	notificationListener := OpenListener("notifications")

	ticker := time.NewTicker(1 * time.Second)

	defer CloseListener(channelName, eventListener)
	defer CloseListener("notifications", notificationListener)

	c.Stream(func(w io.Writer) bool {
		select {
		case event := <-eventListener:
			c.SSEvent("message", event)
		case event := <-notificationListener:
			c.SSEvent("notification", event)
		case <-ticker.C:
			c.SSEvent("ping", nil)
		}

		return true
	})
}
