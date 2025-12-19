package bootstrap

import (
	"github.com/Iversy/unified-message-hub/internal/services/processors/message_processor"
)

func InitMessageProcessor(messageService messageService) *messageprocessor.MessageProcessor {
	return messageprocessor.NewMessageProcessor(messageService)
}
