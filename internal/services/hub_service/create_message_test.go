package hubservice

import (
	"context"
	"errors"
	"testing"

	"github.com/Iversy/unified-message-hub/internal/models"
	"github.com/Iversy/unified-message-hub/internal/services/hub_service/mocks"
	"github.com/stretchr/testify/suite"
	"gotest.tools/v3/assert"
)

type MessageServiceSuite struct {
	suite.Suite
	ctx            context.Context
	messageStorage *mocks.HubStorage
	messageService *HubService
}

func (m *MessageServiceSuite) SetupTest() {
	m.messageStorage = mocks.NewHubStorage(m.T())
	m.ctx = context.Background()
	m.messageService = NewHubService(m.ctx, m.messageStorage)
}

func (m *MessageServiceSuite) TestCreateSuccess() {
	messages := []*models.Message{
		{
			Client:    models.VK,
			Sender:    "oleg",
			ChatId:    3,
			Text:      "biliberda",
			Timestamp: "2024-01-15T10:20:00Z",
		},
	}
	m.messageStorage.EXPECT().CreateMessage(m.ctx, messages).Return(nil)

	err := m.messageStorage.CreateMessage(m.ctx, messages)

	assert.NilError(m.T(), err)

}

func (m *MessageServiceSuite) TestCreateStorageError() {
	messages := []*models.Message{
		{
			Client:    models.VK,
			Sender:    "oleg",
			ChatId:    3,
			Text:      "biliberda",
			Timestamp: "2024-01-15T10:20:00Z",
		},
	}
	wantErr := errors.New("error")

	m.messageStorage.EXPECT().CreateMessage(m.ctx, messages).Return(wantErr)

	err := m.messageStorage.CreateMessage(m.ctx, messages)

	assert.ErrorIs(m.T(), err, wantErr)

}

func (m *MessageServiceSuite) TestCreateEmptyText() {
	messages := []*models.Message{
		{
			Client:    models.VK,
			Sender:    "oleg",
			ChatId:    3,
			Text:      "",
			Timestamp: "2024-01-15T10:20:00Z",
		},
	}

	err := m.messageService.CreateMessage(m.ctx, messages)
	assert.Check(m.T(), err != nil)

}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(MessageServiceSuite))
}
