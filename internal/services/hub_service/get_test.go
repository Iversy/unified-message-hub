package hubservice

import (
	"github.com/Iversy/unified-message-hub/internal/models"
	"github.com/stretchr/testify/assert"
)

func (h *HubServiceSuite) TestGetActiveRoutesBySourceChatID() {
	wantRoutes := []*models.Route{
		{
			ID:           1,
			Name:         "oleg",
			SourceChatID: 666,
			ReceiverID:   111,
			Keywords:     []string{},
			IsActive:     true,
		},
	}
	h.hubStorage.On("GetActiveRoutesBySourceChatID", h.ctx, 666).
		Return(wantRoutes, nil).
		Once()
	gotRoutes, gotErr := h.hubService.GetActiveRoutesBySourceChatID(h.ctx, 666)

	assert.Equal(h.T(), wantRoutes, gotRoutes)
	assert.Nil(h.T(), gotErr)
}
