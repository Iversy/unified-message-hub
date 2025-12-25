package hubprocessor

import (
	"context"
	"fmt"
	"log"

	"github.com/Iversy/unified-message-hub/internal/models"
)

func (p *HubProcessor) HandleMessage(ctx context.Context, message *models.Message) error {

	err := p.hubService.CreateMessage(ctx, []*models.Message{message})
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
	routes, err := p.hubService.GetActiveRoutesBySourceChatID(ctx, message.ChatId) // TODO: filter receivers by keywords
	log.Println("Handling Message", message.Timestamp)
	if err != nil {
		return err
	}
	//for _, route := range routes {
	//	log.Println("Route", route)
	//	err = utils.InspectStruct(*route)
	//	if err != nil {
	//		return err
	//	}
	//}
	return p.vkService.SendMessageMulti(routes, msg)
}
