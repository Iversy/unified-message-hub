package bootstrap

import (
	"context"

	hubservice "github.com/Iversy/unified-message-hub/internal/services/hub_service"
	"github.com/Iversy/unified-message-hub/internal/storage/pgstorage"
)

func InitHubService(storage *pgstorage.PGstorage) *hubservice.HubService {
	return hubservice.NewHubService(context.Background(), storage)
}
