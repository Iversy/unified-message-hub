package hubservice

import (
	"context"

	"github.com/Iversy/unified-message-hub/internal/models"
)

func (h *HubService) GetActiveRoutesBySourceChatID(ctx context.Context, chatID int) ([]*models.Route, error) {
	return h.hubStorage.GetActiveRoutesBySourceChatID(ctx, chatID)
}
