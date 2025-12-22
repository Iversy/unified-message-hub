package bootstrap

import (
	"context"
	"encoding/json"
	"fmt"

	messageconsumer "github.com/Iversy/unified-message-hub/internal/consumer/message_consumer"
	"github.com/Iversy/unified-message-hub/internal/models"
	"github.com/Iversy/unified-message-hub/internal/producer"
	"github.com/gin-gonic/gin"
)

type messageProducer interface {
	ProduceMessage(ctx context.Context, key, value []byte, headers map[string]string) error
}

type messageCreateProducer struct {
	messageProducer messageProducer
}

func (mp *messageCreateProducer) Handle(message *models.Message) error {
	key := []byte(fmt.Sprintf("%v%v", message.ChatId, message.Timestamp))
	value, err := json.Marshal(message)
	if err != nil {
		return err
	}

	headers := map[string]string{}
	mp.messageProducer.ProduceMessage(context.Background(), key, value, headers)

	return nil
}

func NewMessageCreateProducer(messageProducer messageProducer) *messageCreateProducer {
	return &messageCreateProducer{
		messageProducer: messageProducer,
	}
}

func AppRun(producer *producer.KafkaProducer, consumer *messageconsumer.MessageCreateConsumer) {
	go consumer.Consume(context.Background())
	mcp := NewMessageCreateProducer(producer)
	go func() {

		if err := runAPIserver(mcp); err != nil {
			panic(err)
		}
	}()

}

func (mcp *messageCreateProducer) postMessage(c *gin.Context) {
	var newMessage models.Message = *models.NewMessage()
	if err := c.BindJSON(&newMessage); err != nil {
		return
	}
	mcp.Handle(&newMessage)
}

func runAPIserver(mcp *messageCreateProducer) error {
	router := gin.Default()
	router.POST("/api/message/send", mcp.postMessage)

	err := router.Run(":8080")
	if err != nil {
		return err
	}
	return nil
}
