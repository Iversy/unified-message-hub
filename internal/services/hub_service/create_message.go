package hubservice

import (
	"context"

	"github.com/Iversy/unified-message-hub/internal/models"
)

func (h *HubService) CreateMessage(ctx context.Context, messageInfos []*models.Message) error {
	return h.hubStorage.CreateMessage(ctx, messageInfos)
}
