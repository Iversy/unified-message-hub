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
	messageProducer := bootstrap.InitMessageProducer(cfg)
	routeProducer := bootstrap.InitRouteProducer(cfg)
	defer messageProducer.Close()
	defer routeProducer.Close()

	messageService := bootstrap.InitHubService(storage)
	hubProcessor := bootstrap.InitMessageProcessor(messageService, vkService)
	messageConsumer := bootstrap.InitMessageCreateConsumer(cfg, hubProcessor)
	routeConsumer := bootstrap.InitRouteCreateConsumer(cfg, hubProcessor)
	serviceAPI := bootstrap.InitHubServiceAPI(messageService, messageProducer, routeProducer)

	bootstrap.AppRun(serviceAPI, messageConsumer, routeConsumer)
}
