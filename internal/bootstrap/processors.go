package bootstrap

import (
	"fmt"

	"github.com/Iversy/unified-message-hub/config"
	"github.com/Iversy/unified-message-hub/internal/producer"
	messageservice "github.com/Iversy/unified-message-hub/internal/services/message_service"
	messageprocessor "github.com/Iversy/unified-message-hub/internal/services/processors/message_processor"
)

func InitMessageProducer(cfg *config.Config) *producer.KafkaProducer {
	broker := fmt.Sprintf("%v:%v", cfg.Kafka.Host, cfg.Kafka.Port)
	return producer.NewKafkaProducer(broker, cfg.Kafka.MessageCreateTopicName)
}

func InitMessageProcessor(messageService *messageservice.MessageService) *messageprocessor.MessageProcessor {
	return messageprocessor.NewMessageProcessor(messageService)
}
