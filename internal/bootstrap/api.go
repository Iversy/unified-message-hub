package bootstrap

import (
	"github.com/Iversy/unified-message-hub/internal/api"
	"github.com/Iversy/unified-message-hub/internal/producer"
	hubservice "github.com/Iversy/unified-message-hub/internal/services/hub_service"
)

func InitHubServiceAPI(hubService *hubservice.HubService, messageProducer, routeProducer *producer.KafkaProducer) *api.HubServiceAPI {
	return api.NewHubServiceAPI(hubService, messageProducer, routeProducer)
}
