package messageprocessor

import (
	"context"

	"github.com/Iversy/unified-message-hub/internal/models"
)

type messageService interface {
	CreateMessage(ctx context.Context, studentsInfos []*models.Message) error
}

type vkService interface {
	SendBroadcast(message string) error
}

type MessageProcessor struct {
	messageService messageService
	vkService      vkService
}

func NewMessageProcessor(messageService messageService, vkService vkService) *MessageProcessor {
	return &MessageProcessor{
		messageService: messageService,
		vkService:      vkService,
	}
}
