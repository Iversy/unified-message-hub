package bootstrap

import (
	"github.com/Iversy/unified-message-hub/internal/api"
	"github.com/Iversy/unified-message-hub/internal/producer"
	messageservice "github.com/Iversy/unified-message-hub/internal/services/message_service"
)

func InitMessageServiceAPI(messageService *messageservice.MessageService, messageProducer *producer.KafkaProducer) *api.MessageServiceAPI {
	return api.NewMessageServiceAPI(messageService, messageProducer)
}
