package messageprocessor

import (
	"context"

	"github.com/Iversy/unified-message-hub/internal/models"
)

func (p *MessageProcessor) Handle(ctx context.Context, message *models.Message) error {
	return p.messageService.CreateMessage(ctx, []*models.Message{message})
}
