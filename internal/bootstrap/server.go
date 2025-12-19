package bootstrap

import (
	"context"

	messageconsumer "github.com/Iversy/unified-message-hub/internal/consumer/message_consumer"
	"github.com/Iversy/unified-message-hub/internal/models"
	"github.com/gin-gonic/gin"
)

func AppRun(consumer messageconsumer.MessageCreateConsumer) {
	go consumer.Consume(context.Background())
	go func() {
		err := runAPIserver()
		if err != nil {
			panic(err)
		}
	}()

}

func postMessage(c *gin.Context) {
	var newMessage models.Message = *models.NewMessage()
	if err := c.BindJSON(&newMessage); err != nil {
		return
	}
	// Handle(&newMessage)
}

func runAPIserver() error {
	router := gin.Default()
	router.POST("/api/message/send", postMessage)

	err := router.Run(":8080")
	if err != nil {
		return err
	}
	return nil
}
