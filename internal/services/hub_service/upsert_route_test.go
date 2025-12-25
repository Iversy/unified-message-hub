package hubservice

import (
	"github.com/Iversy/unified-message-hub/internal/models"
	"github.com/stretchr/testify/assert"
)

func (h *HubServiceSuite) TestUpsertRoute() {
	routes := []*models.Route{
		{
			ID:           1,
			Name:         "oleg",
			SourceChatID: 666,
			ReceiverID:   111,
			Keywords:     []string{},
			IsActive:     true,
		},
	}
	h.hubStorage.On("UpsertRoute", h.ctx, routes).
		Return(nil).
		Once()
	gotErr := h.hubService.UpsertRoute(h.ctx, routes)

	assert.Nil(h.T(), gotErr)
}
