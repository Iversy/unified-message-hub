package producer

import (
	"fmt"
	"github.com/Iversy/unified-message-hub/config"
	"log"
)

type AppContext struct {
	KafkaProducer *KafkaProducer
}

func NewAppContext(topic string, cfg *config.Config) *AppContext {
	broker := fmt.Sprintf("%s:%s", cfg.Kafka.Host, cfg.Kafka.Port)
	kafkaProducer := NewKafkaProducer(broker, topic)

	return &AppContext{
		KafkaProducer: kafkaProducer,
	}
}

func (ac *AppContext) Shutdown() {
	if err := ac.KafkaProducer.Close(); err != nil {
		log.Printf("Error closing Kafka producer: %v", err)
	}
	log.Println("Application shutdown complete")
}
