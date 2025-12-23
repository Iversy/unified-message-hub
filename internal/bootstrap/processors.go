package bootstrap

import (
	messageservice "github.com/Iversy/unified-message-hub/internal/services/message_service"
	messageprocessor "github.com/Iversy/unified-message-hub/internal/services/processors/message_processor"
	vkservice "github.com/Iversy/unified-message-hub/internal/services/vk_service"
)

func InitMessageProcessor(messageService *messageservice.MessageService, vkService *vkservice.VKService) *messageprocessor.MessageProcessor {
	return messageprocessor.NewMessageProcessor(messageService, vkService)
}
