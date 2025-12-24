package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Iversy/unified-message-hub/internal/models"
	"github.com/Iversy/unified-message-hub/internal/utils"
)

func (mapi *HubServiceAPI) HandleMessage(message *models.Message) error {
	key := []byte(fmt.Sprintf("%v%v", message.ChatId, message.Timestamp))
	value, err := json.Marshal(message)
	if err != nil {
		return err
	}

	headers := map[string]string{}
	mapi.messageProducer.ProduceMessage(context.Background(), key, value, headers)

	return nil
}

func (mapi *HubServiceAPI) HandleRoute(route *models.Route) error {
	key := []byte(fmt.Sprintf("%v%v", route.Name, route.SourceChatID))
	value, err := json.Marshal(route)
	if err != nil {
		return err
	}
	utils.InspectStruct(route)
	headers := map[string]string{}
	mapi.routeProducer.ProduceMessage(context.Background(), key, value, headers)

	return nil
}
