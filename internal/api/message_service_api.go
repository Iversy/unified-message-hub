package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Iversy/unified-message-hub/internal/models"
)

type messageProducer interface {
	ProduceMessage(ctx context.Context, key, value []byte, headers map[string]string) error
}

type messageService interface {
	CreateMessage(ctx context.Context, messageInfos []*models.Message) error
}

type MessageServiceAPI struct {
	messageService  messageService
	messageProducer messageProducer
}

func (mapi *MessageServiceAPI) Handle(message *models.Message) error {
	key := []byte(fmt.Sprintf("%v%v", message.ChatId, message.Timestamp))
	value, err := json.Marshal(message)
	if err != nil {
		return err
	}

	headers := map[string]string{}
	mapi.messageProducer.ProduceMessage(context.Background(), key, value, headers)

	return nil
}

func NewMessageServiceAPI(messageService messageService, messageProducer messageProducer) *MessageServiceAPI {
	return &MessageServiceAPI{
		messageService:  messageService,
		messageProducer: messageProducer,
	}
}
