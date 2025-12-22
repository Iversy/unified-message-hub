package bootstrap

import (
	"context"

	"github.com/Iversy/unified-message-hub/internal/producer"
	messageservice "github.com/Iversy/unified-message-hub/internal/services/message_service"
	"github.com/Iversy/unified-message-hub/internal/storage/pgstorage"
)

func InitMessageService(storage *pgstorage.PGstorage, messageProducer *producer.KafkaProducer) *messageservice.MessageService {
	return messageservice.NewMessageService(context.Background(), storage, messageProducer)
}
