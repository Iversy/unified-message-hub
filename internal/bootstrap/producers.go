package bootstrap

import (
	"fmt"

	"github.com/Iversy/unified-message-hub/config"
	"github.com/Iversy/unified-message-hub/internal/producer"
)

func InitMessageProducer(cfg *config.Config) *producer.KafkaProducer {
	broker := fmt.Sprintf("%v:%v", cfg.Kafka.Host, cfg.Kafka.Port)
	return producer.NewKafkaProducer(broker, cfg.Kafka.MessageCreateTopicName)
}
