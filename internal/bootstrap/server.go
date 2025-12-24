package bootstrap

import (
	"context"

	"github.com/Iversy/unified-message-hub/internal/api"
	messageconsumer "github.com/Iversy/unified-message-hub/internal/consumer/message_consumer"
	routeconsumer "github.com/Iversy/unified-message-hub/internal/consumer/route_consumer"
	"github.com/gin-gonic/gin"
)

func AppRun(service_api *api.HubServiceAPI, messageConsumer *messageconsumer.MessageCreateConsumer, routeConsumer *routeconsumer.RouteCreateConsumer) {
	go messageConsumer.Consume(context.Background())
	go routeConsumer.Consume(context.Background())
	if err := runAPIserver(service_api); err != nil {
		panic(err)
	}
}

func runAPIserver(mapi *api.HubServiceAPI) error {
	router := gin.Default()
	router.POST("/api/send/message", mapi.PostMessage)
	router.POST("/api/send/route", mapi.PostRoute)

	err := router.Run(":8080")
	if err != nil {
		return err
	}
	return nil
}
