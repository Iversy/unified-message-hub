package bootstrap

import (
	hubservice "github.com/Iversy/unified-message-hub/internal/services/hub_service"
	hubprocessor "github.com/Iversy/unified-message-hub/internal/services/processors/hub_processor"
	vkservice "github.com/Iversy/unified-message-hub/internal/services/vk_service"
)

func InitMessageProcessor(hubService *hubservice.HubService, vkService *vkservice.VKService) *hubprocessor.HubProcessor {
	return hubprocessor.NewHubProcessor(hubService, vkService)
}
