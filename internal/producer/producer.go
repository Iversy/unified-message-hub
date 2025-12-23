package producer

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
	topic  string
}

func NewKafkaProducer(broker string, topic string) *KafkaProducer {
	writer := &kafka.Writer{
		Addr:         kafka.TCP(broker),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		MaxAttempts:  3,
		BatchSize:    100,
		BatchBytes:   1048576,
		BatchTimeout: 10 * time.Millisecond,
		RequiredAcks: kafka.RequireAll,
		Async:        false,
		Compression:  kafka.Snappy,
	}
	return &KafkaProducer{
		writer: writer,
		topic:  topic,
	}

}

func (kp *KafkaProducer) ProduceMessage(ctx context.Context, key, value []byte, headers map[string]string) error {
	msg := kafka.Message{
		Key:   key,
		Value: value,
		Time:  time.Now(),
	}

	for k, v := range headers {
		msg.Headers = append(msg.Headers, kafka.Header{
			Key:   k,
			Value: []byte(v),
		})
	}
	return kp.writer.WriteMessages(ctx, msg)

}

func (kp *KafkaProducer) Close() error {
	return kp.writer.Close()
}
