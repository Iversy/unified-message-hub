package hubservice

import (
	"errors"

	"github.com/Iversy/unified-message-hub/internal/models"
	"github.com/stretchr/testify/assert"
)

func (m *HubServiceSuite) TestCreateStorageError() {
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

	m.hubStorage.EXPECT().CreateMessage(m.ctx, messages).Return(wantErr)

	err := m.hubStorage.CreateMessage(m.ctx, messages)

	assert.ErrorIs(m.T(), err, wantErr)

}

func (h *HubServiceSuite) TestCreateMessageInvalidClientError() {
	messages := []*models.Message{
		{
			Client:    3,
			Sender:    "oleg",
			ChatId:    3,
			Text:      "biliberda",
			Timestamp: "2024-01-15T10:20:00Z",
		},
	}
	wantErr := "ошибка соответствия платформы, введена несуществующая платформа"
	gotErr := h.hubService.CreateMessage(h.ctx, messages)

	assert.EqualError(h.T(), gotErr, wantErr)
}

func (h *HubServiceSuite) TestCreateMessageSuccess() {
	messages := []*models.Message{
		{
			Client:    models.Telegram,
			Sender:    "oleg",
			ChatId:    3,
			Text:      "biliberda",
			Timestamp: "2024-01-15T10:20:00Z",
		},
		{
			Client:    models.Telegram,
			Sender:    "oleg",
			ChatId:    3,
			Text:      "biliberda",
			Timestamp: "2024-01-15T10:20:00Z",
		},
	}

	h.hubStorage.On("CreateMessage", h.ctx, messages).
		Return(nil).
		Once()

	err := h.hubService.CreateMessage(h.ctx, messages)

	assert.Nil(h.T(), err)
}

func (h *HubServiceSuite) TestCreateMessageFutureError() {
	messages := []*models.Message{
		{
			Client:    models.Telegram,
			Sender:    "oleg",
			ChatId:    3,
			Text:      "biliberda",
			Timestamp: "3000-01-15T10:20:00Z",
		},
	}
	wantErr := "ошибка привет из будущего"

	err := h.hubService.CreateMessage(h.ctx, messages)

	assert.EqualError(h.T(), err, wantErr)
}

func (h *HubServiceSuite) TestCreateMessageParsingError() {
	messages := []*models.Message{
		{
			Client:    models.Telegram,
			Sender:    "oleg",
			ChatId:    3,
			Text:      "biliberda",
			Timestamp: "300000-01-15T10:20:00",
		},
	}
	wantErr := "parsing error"

	err := h.hubService.CreateMessage(h.ctx, messages)

	assert.EqualError(h.T(), err, wantErr)
}
