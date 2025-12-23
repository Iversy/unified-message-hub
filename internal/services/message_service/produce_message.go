package messageservice

import (
	"context"
)

func (s *MessageService) ProduceMessage(ctx context.Context, key, value []byte, headers map[string]string) error {
	return s.messageProducer.ProduceMessage(ctx, key, value, headers)
}
