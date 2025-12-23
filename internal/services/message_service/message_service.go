package messageservice

import (
	"context"

	"github.com/Iversy/unified-message-hub/internal/models"
)

type MessageStorage interface {
	CreateMessage(ctx context.Context, messageInfos []*models.Message) error
}

type MessageProducer interface {
	ProduceMessage(ctx context.Context, key, value []byte, headers map[string]string) error
}

type MessageService struct {
	messageStorage  MessageStorage
	messageProducer MessageProducer
}

func NewMessageService(ctx context.Context, messageStorage MessageStorage, messageProducer MessageProducer) *MessageService {
	return &MessageService{
		messageStorage:  messageStorage,
		messageProducer: messageProducer,
	}
}
