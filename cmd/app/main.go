package main

import (
	"fmt"

	"github.com/Iversy/unified-message-hub/config"
	"github.com/Iversy/unified-message-hub/internal/bootstrap"
)

func main() {

	cfg, err := config.LoadConfig("config_norm.yaml")
	if err != nil {
		panic(fmt.Sprintf("ошибка парсинга конфига, %v", err))
	}

	storage := bootstrap.InitPGStorage(cfg)
	kafkaProducer := bootstrap.InitMessageProducer(cfg)
	messageService := bootstrap.InitMessageService(storage, kafkaProducer)
	messageProcessor := bootstrap.InitMessageProcessor(messageService)
	kafkaConsumer := bootstrap.InitMessageCreateConsumer(cfg, messageProcessor)
	serviceAPI := bootstrap.InitMessageServiceAPI(messageService, kafkaProducer)
	bootstrap.AppRun(serviceAPI, kafkaConsumer)
	//bootstrap.AppRun()
	//studentService := bootstrap.InitStudentService(studentsStorage, cfg)
	//studentsInfoProcessor := bootstrap.InitStudentsInfoProcessor(studentService)
	//studentsinfoupsertconsumer := bootstrap.InitStudentInfoUpsertConsumer(cfg, studentsInfoProcessor)
	//studentsApi := bootstrap.InitStudentServiceAPI(studentService)
	//
	//bootstrap.AppRun(*studentsApi, studentsinfoupsertconsumer)
}
