package messageconsumer

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/Iversy/unified-message-hub/internal/models"
	"github.com/segmentio/kafka-go"
)

func (c *MessageCreateConsumer) Consume(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:           c.kafkaBroker,
		Partition:         0,
		Topic:             c.topicName,
		HeartbeatInterval: 3 * time.Second,
		SessionTimeout:    30 * time.Second,
	})
	defer r.Close()

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			slog.Error("MessageInfoUpsertConsumer.consume error", "error", err.Error())
		}
		var message *models.Message
		err = json.Unmarshal(msg.Value, &message)
		if err != nil {
			slog.Error("parse", "error", err)
			continue
		}
		err = c.messageProcessor.Handle(ctx, message)
		if err != nil {
			slog.Error("Handle", "error", err)
		}
	}

}
