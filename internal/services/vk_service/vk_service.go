package vkservice

import (
	"sync"

	"github.com/Iversy/unified-message-hub/config"
	"github.com/SevereCloud/vksdk/v3/api"
	"github.com/SevereCloud/vksdk/v3/longpoll-bot"
)

type VKService struct {
	VK          *api.VK
	mutex       *sync.RWMutex
	lp          *longpoll.LongPoll
	config      *config.VK
	subscribers map[int]bool
}

func NewVKService(vk *api.VK, mutex *sync.RWMutex, lp *longpoll.LongPoll, cfg *config.VK, subscribers map[int]bool) *VKService {
	return &VKService{
		VK:          vk,
		mutex:       mutex,
		lp:          lp,
		config:      cfg,
		subscribers: subscribers,
	}
}
