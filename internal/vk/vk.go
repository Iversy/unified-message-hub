package main

import (
	"fmt"
	"log"

	"github.com/SevereCloud/vksdk/v3/api"
)

func main() {
	// Access token for user.
	token := ""
	fmt.Println(token)
	vk := api.NewVK(token)

	parameters := api.Params{
		"chat_id": 75,
	}

	messages, err := vk.MessagesGetChat(parameters)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages.Title)

	// // get information about the group
	// group, err := vk.GroupsGetByID(nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Initializing Long Poll
	// lp, err := longpoll.NewLongPoll(vk, group.Groups[0].ID)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // New message event
	// lp.MessageNew(func(_ context.Context, obj events.MessageNewObject) {
	// 	log.Printf("%d: %s", obj.Message.PeerID, obj.Message.Text)

	// 	if obj.Message.Text == "ping" {
	// 		b := params.NewMessagesSendBuilder()
	// 		b.Message("pong")
	// 		b.RandomID(0)
	// 		b.PeerID(obj.Message.PeerID)

	// 		_, err := vk.MessagesSend(b.Params)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 	}
	// })

	// // Run Bots Long Poll
	// log.Println("Start Long Poll")
	// if err := lp.Run(); err != nil {
	// 	log.Fatal(err)
	// }

}
