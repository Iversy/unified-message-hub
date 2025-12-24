package vkservice

import (
	"fmt"
	"log"

	"github.com/SevereCloud/vksdk/v3/api/params"
)

func (v *VKService) updateSubscribersCache() {
	b := params.NewGroupsGetMembersBuilder()
	b.GroupID(fmt.Sprint(v.config.GroupID))
	b.Filter("unsure")

	resp, err := v.VK.GroupsGetMembers(b.Params)
	if err != nil {
		log.Printf("Ошибка получения подписчиков: %v", err)
		return
	}

	v.mutex.Lock()
	for _, userID := range resp.Items {
		v.subscribers[userID] = true
	}
	v.mutex.Unlock()

	log.Printf("Загружено %d подписчиков в кэш", len(resp.Items))
}
