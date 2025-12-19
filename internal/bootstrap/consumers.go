package bootstrap

import (
	"fmt"

	"github.com/Iversy/unified-message-hub/config"
	messageconsumer "github.com/Iversy/unified-message-hub/internal/consumer/message_consumer"
	messageservice "github.com/Iversy/unified-message-hub/internal/services/message_service"
)

func InitMessageCreateConsumer(cfg *config.Config, messageService *messageservice.MessageService) *messageconsumer.MessageCreateConsumer {
	kafkaBrokers := []string{fmt.Sprintf("%v:%v", cfg.Kafka.Host, cfg.Kafka.Port)}
	return messageconsumer.NewMessageCreateConsumer(messageService, kafkaBrokers, cfg.Kafka.MessageCreateTopicName)
}
