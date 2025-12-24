package routeconsumer

import (
	"context"

	"github.com/Iversy/unified-message-hub/internal/models"
)

type routeProcessor interface {
	HandleRoute(ctx context.Context, messageInfo *models.Route) error
}

type RouteCreateConsumer struct {
	routeProcessor routeProcessor
	kafkaBroker    []string
	topicName      string
}

func NewRouteCreateConsumer(hubProcessor routeProcessor, kafkaBroker []string, topicName string) *RouteCreateConsumer {
	return &RouteCreateConsumer{
		routeProcessor: hubProcessor,
		kafkaBroker:    kafkaBroker,
		topicName:      topicName,
	}
}
