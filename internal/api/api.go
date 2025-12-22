package api

import (
	"fmt"

	"github.com/Iversy/unified-message-hub/internal/models"
	"github.com/gin-gonic/gin"
)

func (mapi *MessageServiceAPI) PostMessage(c *gin.Context) {
	var newMessage models.Message = *models.NewMessage()
	if err := c.BindJSON(&newMessage); err != nil {
		return
	}
	fmt.Printf("----------%v\n", newMessage)
	mapi.Handle(&newMessage)
}
