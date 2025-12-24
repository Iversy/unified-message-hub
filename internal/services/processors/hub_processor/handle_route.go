package hubprocessor

import (
	"context"

	"github.com/Iversy/unified-message-hub/internal/models"
)

func (p *HubProcessor) HandleRoute(ctx context.Context, route *models.Route) error {

	err := p.hubService.UpsertRoute(ctx, []*models.Route{route})
	if err != nil {
		return err
	}
	return nil
}
