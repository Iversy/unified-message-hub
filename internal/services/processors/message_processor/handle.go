package messageprocessor

import (
	"context"
	"fmt"

	"github.com/Iversy/unified-message-hub/internal/models"
)

func (p *MessageProcessor) Handle(ctx context.Context, message *models.Message) error {

	err := p.messageService.CreateMessage(ctx, []*models.Message{message})
	if err != nil {
		return err
	}
	msg := fmt.Sprintf(`
		%v:%v

		%v
		
		%v
		`,
		message.Client, message.Sender, message.Text, message.Timestamp,
	)
	return p.vkService.SendBroadcast(msg)
}
