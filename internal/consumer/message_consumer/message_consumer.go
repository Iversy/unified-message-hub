package messageconsumer

import (
	"context"

	"github.com/Iversy/unified-message-hub/internal/models"
)

type messageProcessor interface {
	Handle(ctx context.Context, messageInfo *models.Message) error
}

type MessageCreateConsumer struct {
	messageProcessor messageProcessor
	kafkaBroker      []string
	topicName        string
}

func NewMessageCreateConsumer(messageProcessor messageProcessor, kafkaBroker []string, topicName string) *MessageCreateConsumer {
	return &MessageCreateConsumer{
		messageProcessor: messageProcessor,
		kafkaBroker:      kafkaBroker,
		topicName:        topicName,
	}
}
