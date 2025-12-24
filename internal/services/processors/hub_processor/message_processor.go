package hubprocessor

import (
	"context"

	"github.com/Iversy/unified-message-hub/internal/models"
)

type hubService interface {
	CreateMessage(ctx context.Context, messageInfos []*models.Message) error
	UpsertRoute(ctx context.Context, routes []*models.Route) error
	GetActiveRoutesBySourceChatID(ctx context.Context, chatID int) ([]*models.Route, error)
}

type vkService interface {
	SendBroadcast(message string) error
	SendMessageMulti(routes []*models.Route, text string) error
}

type HubProcessor struct {
	hubService hubService
	vkService  vkService // мб сделать ещё слой абстракции
}

func NewHubProcessor(hubService hubService, vkService vkService) *HubProcessor {
	return &HubProcessor{
		hubService: hubService,
		vkService:  vkService,
	}
}
