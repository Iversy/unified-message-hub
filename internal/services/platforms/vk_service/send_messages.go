package vkservice

import (
	"fmt"
	"log"
	"time"

	"github.com/SevereCloud/vksdk/v3/api/params"
)

func (v *VKService) SendBroadcast(message string) error {
	v.mutex.RLock()
	userIDs := make([]int, 0, len(v.subscribers))
	for id := range v.subscribers {
		userIDs = append(userIDs, id)
	}
	v.mutex.RUnlock()

	log.Printf("Начинаю рассылку для %d пользователей...", len(userIDs))

	successCount := 0
	for _, userID := range userIDs {

		err := v.sendMessage(userID, message)
		if err != nil {
			log.Printf("Рассылка завершена с ошибкой. Успешно: %d/%d", successCount, len(userIDs))
			return fmt.Errorf("Ошибка отправки пользователю %d: %v", userID, err)
		}
		successCount++

		// VK API лимит 3 запроса/сек
		time.Sleep(v.config.Delay)
	}

	log.Printf("Рассылка завершена. Успешно: %d/%d", successCount, len(userIDs))
	// log.Printf("Ухожу в timeout на %v секунд", successCount, len(userIDs))
	// time.Sleep(v.config.Timeout)
	return nil
}

func (v *VKService) sendMessage(userID int, text string) error {
	b := params.NewMessagesSendBuilder()
	b.Message(text)
	b.RandomID(0)
	b.PeerID(userID)

	_, err := v.VK.MessagesSend(b.Params)
	return err
}
