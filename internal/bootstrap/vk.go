package bootstrap

import (
	"os"
	"sync"

	"github.com/Iversy/unified-message-hub/config"
	vkservice "github.com/Iversy/unified-message-hub/internal/services/platforms/vk_service"
	"github.com/SevereCloud/vksdk/v3/api"
	"github.com/SevereCloud/vksdk/v3/longpoll-bot"
)

func InitVKService(cfg *config.Config) (*vkservice.VKService, error) {
	var mutex sync.RWMutex
	token := os.Getenv("VKTOKEN")
	vk := api.NewVK(token)
	subscribers := make(map[int]bool)
	lp, err := longpoll.NewLongPoll(vk, cfg.VK.GroupID)
	if err != nil {
		return nil, err
	}
	return vkservice.NewVKService(vk, &mutex, lp, &cfg.VK, subscribers), nil
}

func VKListen(vkservice *vkservice.VKService) {
	vkservice.Listen()
}
