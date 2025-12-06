package listener

import (
	"github.com/Iversy/unified-message-hub/internal/models"
)

func Handle(message *models.Message) (models.Message, error) {

	return *models.NewMessage(), nil
}
