package main

import (
	"log"

	"github.com/Iversy/unified-message-hub/config"
	"github.com/Iversy/unified-message-hub/internal/bootstrap"
)

func main() {

	cfg, err := config.LoadConfig("config_norm.yaml")
	if err != nil {
		log.Panicf("ошибка парсинга конфига, %v", err)
	}

	vkService, err := bootstrap.InitVKService(cfg)
	if err != nil {
		log.Printf("Ошибка инициализации вк, %v", err)
	}
	go bootstrap.VKListen(vkService)

	storage := bootstrap.InitPGStorage(cfg)
	kafkaProducer := bootstrap.InitMessageProducer(cfg)
	defer kafkaProducer.Close()

	messageService := bootstrap.InitMessageService(storage)
	messageProcessor := bootstrap.InitMessageProcessor(messageService, vkService)
	kafkaConsumer := bootstrap.InitMessageCreateConsumer(cfg, messageProcessor)
	serviceAPI := bootstrap.InitMessageServiceAPI(messageService, kafkaProducer)

	bootstrap.AppRun(serviceAPI, kafkaConsumer)
}
