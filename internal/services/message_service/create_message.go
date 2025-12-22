package messageservice

import (
	"context"

	"github.com/Iversy/unified-message-hub/internal/models"
)

func (s *MessageService) CreateMessage(ctx context.Context, messageInfos []*models.Message) error {
	return s.messageStorage.CreateMessage(ctx, messageInfos)
}
