package api

import (
	"context"

	"github.com/Iversy/unified-message-hub/internal/models"
)

type messageProducer interface {
	ProduceMessage(ctx context.Context, key, value []byte, headers map[string]string) error
}

type hubService interface {
	CreateMessage(ctx context.Context, messageInfos []*models.Message) error
}
type routeProducer interface {
	ProduceMessage(ctx context.Context, key, value []byte, headers map[string]string) error
}

type HubServiceAPI struct {
	hubService      hubService
	messageProducer messageProducer
	routeProducer   routeProducer
}

func NewHubServiceAPI(hubService hubService, messageProducer messageProducer, routeProducer routeProducer) *HubServiceAPI {
	return &HubServiceAPI{
		hubService:      hubService,
		messageProducer: messageProducer,
		routeProducer:   routeProducer,
	}
}
