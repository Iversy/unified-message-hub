package bootstrap

import (
	"fmt"

	"github.com/Iversy/unified-message-hub/config"
	messageconsumer "github.com/Iversy/unified-message-hub/internal/consumer/message_consumer"
	routeconsumer "github.com/Iversy/unified-message-hub/internal/consumer/route_consumer"
	hubprocessor "github.com/Iversy/unified-message-hub/internal/services/processors/hub_processor"
)

func InitMessageCreateConsumer(cfg *config.Config, hubProccessor *hubprocessor.HubProcessor) *messageconsumer.MessageCreateConsumer {
	kafkaBrokers := []string{fmt.Sprintf("%v:%v", cfg.Kafka.Host, cfg.Kafka.Port)}
	return messageconsumer.NewMessageCreateConsumer(hubProccessor, kafkaBrokers, cfg.Kafka.MessageCreateTopicName)
}

func InitRouteCreateConsumer(cfg *config.Config, hubProccessor *hubprocessor.HubProcessor) *routeconsumer.RouteCreateConsumer {
	kafkaBrokers := []string{fmt.Sprintf("%v:%v", cfg.Kafka.Host, cfg.Kafka.Port)}
	return routeconsumer.NewRouteCreateConsumer(hubProccessor, kafkaBrokers, cfg.Kafka.RouteCreateTopicName)
}
