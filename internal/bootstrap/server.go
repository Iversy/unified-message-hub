package bootstrap

import (
	"github.com/Iversy/unified-message-hub/internal/models"
	"github.com/gin-gonic/gin"
)

func AppRun() {
	
}

func postMessage(c *gin.Context) {
	var newMessage models.Message = *models.NewMessage()
	if err := c.BindJSON(&newMessage); err != nil {
		return
	}
	// Handle(&newMessage)
}

func runAPIserver() {
	router := gin.Default()
	router.POST("/api/message/send", postMessage)
}
