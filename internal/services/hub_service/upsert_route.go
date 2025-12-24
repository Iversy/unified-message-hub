package hubservice

import (
	"context"

	"github.com/Iversy/unified-message-hub/internal/models"
)

func (h *HubService) UpsertRoute(ctx context.Context, routes []*models.Route) error {
	return h.hubStorage.UpsertRoute(ctx, routes)
}
