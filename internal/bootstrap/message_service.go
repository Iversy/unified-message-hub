package bootstrap

import (
	"context"

	messageservice "github.com/Iversy/unified-message-hub/internal/services/message_service"
	"github.com/Iversy/unified-message-hub/internal/storage/pgstorage"
)

func InitMessageService(storage *pgstorage.PGstorage) *messageservice.MessageService {
	return messageservice.NewMessageService(context.Background(), storage)
}
