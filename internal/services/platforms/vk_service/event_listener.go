package vkservice

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/SevereCloud/vksdk/v3/events"
)

func (v *VKService) Listen() {
	v.updateSubscribersCache()
	log.Printf("%v", v.subscribers)

	v.lp.MessageNew(func(ctx context.Context, obj events.MessageNewObject) {
		if obj.Message.Important {
			return
		}

		text := strings.ToLower(strings.TrimSpace(obj.Message.Text))
		fromID := obj.Message.FromID
		log.Printf("%v", text)

		// Команда админа для теста рассылки
		if fromID == v.config.AdminID && strings.HasPrefix(text, "/sendall ") {
			message := strings.TrimPrefix(text, "/sendall ")
			log.Printf("preparing message: %v", message)

			go v.SendBroadcast(message)
			v.sendMessage(fromID, v.config.WelcomeText)
		}
	})

	// новый подписчик
	v.lp.GroupJoin(func(ctx context.Context, obj events.GroupJoinObject) {
		userID := obj.UserID
		log.Printf("Subscribed: %d", userID)

		v.mutex.Lock()
		v.subscribers[userID] = true
		v.mutex.Unlock()

		go func() {
			time.Sleep(1 * time.Second)
			v.sendMessage(userID, v.config.WelcomeText)
		}()
	})

	// отписка
	v.lp.GroupLeave(func(ctx context.Context, obj events.GroupLeaveObject) {
		userID := obj.UserID
		log.Printf("Unsubscribed: %d", userID)

		v.mutex.Lock()
		delete(v.subscribers, userID)
		v.mutex.Unlock()
	})

	log.Println("VK service listening...")
	if err := v.lp.Run(); err != nil {
		log.Fatal("VK service error:", err)
	}
}
