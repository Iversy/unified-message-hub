package messageservice

import (
	"context"

	"github.com/Iversy/unified-message-hub/internal/models"
)

type MessageStorage interface {
	CreateMessage(ctx context.Context, messageInfos []*models.Message) error
}

type MessageService struct {
	messageStorage MessageStorage
}

func NewMessageService(ctx context.Context, messageStorage MessageStorage) *MessageService {
	return &MessageService{
		messageStorage: messageStorage,
	}
}
