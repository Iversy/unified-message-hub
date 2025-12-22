package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Iversy/unified-message-hub/internal/models"
)

type messageProccessor interface {
	ProduceMessage(ctx context.Context, key, value []byte, headers map[string]string) error
}

type MessageHandlerImpl struct {
	messageProccessor messageProccessor
}

func NewMessageHandler(messageProccessor messageProccessor) *MessageHandlerImpl {
	return &MessageHandlerImpl{
		messageProccessor: messageProccessor,
	}
}

func (mh *MessageHandlerImpl) HandleMessage(ctx context.Context, msg *models.Message) error {
	key := fmt.Sprintf("%v%v", msg.ChatId, msg.Timestamp)
	messageJSON, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	err = mh.messageProccessor.ProduceMessage(ctx, []byte(key), messageJSON, nil)
	if err != nil {
		return err
	}
	return nil
}
