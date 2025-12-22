package bootstrap

import (
	"fmt"

	"github.com/Iversy/unified-message-hub/config"
	messageconsumer "github.com/Iversy/unified-message-hub/internal/consumer/message_consumer"
	messageprocessor "github.com/Iversy/unified-message-hub/internal/services/processors/message_processor"
)

func InitMessageCreateConsumer(cfg *config.Config, messageProccessor *messageprocessor.MessageProcessor) *messageconsumer.MessageCreateConsumer {
	kafkaBrokers := []string{fmt.Sprintf("%v:%v", cfg.Kafka.Host, cfg.Kafka.Port)}
	return messageconsumer.NewMessageCreateConsumer(messageProccessor, kafkaBrokers, cfg.Kafka.MessageCreateTopicName)
}
