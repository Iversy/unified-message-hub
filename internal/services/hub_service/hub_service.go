package hubservice

import (
	"context"

	"github.com/Iversy/unified-message-hub/internal/models"
)

type HubStorage interface {
	CreateMessage(ctx context.Context, messageInfos []*models.Message) error
	UpsertRoute(ctx context.Context, routes []*models.Route) error
}

type HubService struct {
	hubStorage HubStorage
}

func NewHubService(ctx context.Context, hubStorage HubStorage) *HubService {
	return &HubService{
		hubStorage: hubStorage,
	}
}
