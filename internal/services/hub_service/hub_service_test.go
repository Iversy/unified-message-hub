package hubservice

import (
	"context"
	"testing"

	"github.com/Iversy/unified-message-hub/internal/services/hub_service/mocks"
	"github.com/stretchr/testify/suite"
)

type HubServiceSuite struct {
	suite.Suite
	ctx        context.Context
	hubStorage *mocks.HubStorage
	hubService *HubService
}

func (m *HubServiceSuite) SetupTest() {
	m.hubStorage = mocks.NewHubStorage(m.T())
	m.ctx = context.Background()
	m.hubService = NewHubService(m.ctx, m.hubStorage)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(HubServiceSuite))
}
