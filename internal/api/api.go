package api

import (
	"github.com/Iversy/unified-message-hub/internal/models"
	"github.com/gin-gonic/gin"
)

func (mapi *HubServiceAPI) PostMessage(c *gin.Context) {
	var newMessage models.Message = *models.NewMessage()
	if err := c.BindJSON(&newMessage); err != nil {
		return
	}
	mapi.HandleMessage(&newMessage)
}

func (mapi *HubServiceAPI) PostRoute(c *gin.Context) {
	var newRoute models.Route = *models.NewRoute()
	if err := c.BindJSON(&newRoute); err != nil {
		return
	}
	mapi.HandleRoute(&newRoute)
}
