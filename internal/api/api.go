package api

import (
	"github.com/Iversy/unified-message-hub/internal/models"
	"github.com/gin-gonic/gin"
)

func (mapi *MessageServiceAPI) PostMessage(c *gin.Context) {
	var newMessage models.Message = *models.NewMessage()
	if err := c.BindJSON(&newMessage); err != nil {
		return
	}
	mapi.Handle(&newMessage)
}
