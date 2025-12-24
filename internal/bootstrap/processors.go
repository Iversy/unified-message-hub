package bootstrap

import (
	hubservice "github.com/Iversy/unified-message-hub/internal/services/hub_service"
	vkservice "github.com/Iversy/unified-message-hub/internal/services/platforms/vk_service"
	hubprocessor "github.com/Iversy/unified-message-hub/internal/services/processors/hub_processor"
)

func InitMessageProcessor(hubService *hubservice.HubService, vkService *vkservice.VKService) *hubprocessor.HubProcessor {
	return hubprocessor.NewHubProcessor(hubService, vkService)
}
