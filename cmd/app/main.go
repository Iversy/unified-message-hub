package main

import (
	"fmt"
	"os"

	"github.com/Iversy/unified-message-hub/config"
	"github.com/Iversy/unified-message-hub/internal/bootstrap"
)

func main() {

	cfg, err := config.LoadConfig(os.Getenv("configPath"))
	if err != nil {
		panic(fmt.Sprintf("ошибка парсинга конфига, %v", err))
	}

	bootstrap.InitPGStorage(cfg)
	bootstrap.InitMessageCreateConsumer(cfg)
	//bootstrap.AppRun()
	//studentService := bootstrap.InitStudentService(studentsStorage, cfg)
	//studentsInfoProcessor := bootstrap.InitStudentsInfoProcessor(studentService)
	//studentsinfoupsertconsumer := bootstrap.InitStudentInfoUpsertConsumer(cfg, studentsInfoProcessor)
	//studentsApi := bootstrap.InitStudentServiceAPI(studentService)
	//
	//bootstrap.AppRun(*studentsApi, studentsinfoupsertconsumer)
}
