package messageprocessor

import (
	"context"

	"github.com/Iversy/unified-message-hub/internal/models"
)

type messageService interface {
	CreateMessage(ctx context.Context, studentsInfos []*models.Message) error
}

type MessageProcessor struct {
	messageService messageService
}

func NewMessageProcessor(messageService messageService) *MessageProcessor {
	return &MessageProcessor{
		messageService: messageService,
	}
}
