package hubservice

import (
	"context"
	"fmt"
	"time"

	"github.com/Iversy/unified-message-hub/internal/models"
)

func (h *HubService) CreateMessage(ctx context.Context, messageInfos []*models.Message) error {
	if err := validateMessageMulti(messageInfos); err != nil {
		return err
	}
	return h.hubStorage.CreateMessage(ctx, messageInfos)
}

func validateMessageMulti(messageInfos []*models.Message) error {
	for _, message := range messageInfos {
		if err := validateMessage(message); err != nil {
			return err
		}
	}
	return nil
}

func validateMessage(message *models.Message) error {
	if message.Client > 2 || message.Client < 0 {
		return fmt.Errorf("ошибка соответствия платформы, введена несуществующая платформа")
	}
	timeParsed, err := time.Parse(time.RFC3339, message.Timestamp)
	if err != nil {
		return fmt.Errorf("parsing error")
	}
	if time.Now().Before(timeParsed) {
		return fmt.Errorf("ошибка привет из будущего")
	}
	return nil
}
