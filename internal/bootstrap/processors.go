package bootstrap

import (
	messageservice "github.com/Iversy/unified-message-hub/internal/services/message_service"
	messageprocessor "github.com/Iversy/unified-message-hub/internal/services/processors/message_processor"
)

func InitMessageProcessor(messageService *messageservice.MessageService) *messageprocessor.MessageProcessor {
	return messageprocessor.NewMessageProcessor(messageService)
}
