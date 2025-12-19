package messageservice

import (
	"context"

	"github.com/Iversy/unified-message-hub/internal/models"
)

type MessageStorage interface {
	//GetStudentInfoByIDs(ctx context.Context, IDs []uint64) ([]*models.Message, error)
	CreateMessage(ctx context.Context, studentInfos []*models.Message) error
}

type MessageService struct {
	messageStorage MessageStorage
}

func NewStudentService(ctx context.Context, messageStorage MessageStorage) *MessageService {
	return &MessageService{
		messageStorage: messageStorage,
	}
}
