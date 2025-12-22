package bootstrap

import (
	"context"

	"github.com/Iversy/unified-message-hub/internal/api"
	messageconsumer "github.com/Iversy/unified-message-hub/internal/consumer/message_consumer"
	"github.com/gin-gonic/gin"
)

func AppRun(service_api *api.MessageServiceAPI, consumer *messageconsumer.MessageCreateConsumer) {
	go consumer.Consume(context.Background())
	if err := runAPIserver(service_api); err != nil {
		panic(err)
	}
	// go func() {

	// }()
}

func runAPIserver(mapi *api.MessageServiceAPI) error {
	router := gin.Default()
	router.POST("/api/send", mapi.PostMessage)

	err := router.Run(":8080")
	if err != nil {
		return err
	}
	return nil
}
