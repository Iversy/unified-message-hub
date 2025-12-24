package routeconsumer

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/Iversy/unified-message-hub/internal/models"
	"github.com/segmentio/kafka-go"
)

func (c *RouteCreateConsumer) Consume(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:           c.kafkaBroker,
		GroupID:           "Route_group",
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
		var route *models.Route
		err = json.Unmarshal(msg.Value, &route)
		if err != nil {
			slog.Error("parse", "error", err)
			continue
		}
		err = c.routeProcessor.HandleRoute(ctx, route)
		if err != nil {
			slog.Error("Handle", "error", err)
		}
	}

}
