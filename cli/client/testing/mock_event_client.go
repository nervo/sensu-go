package testing

import "github.com/sensu/sensu-go/types"

// FetchEvent for use with mock lib
func (c *MockClient) FetchEvent(entity, check string) (*types.Event, error) {
	args := c.Called(entity, check)
	return args.Get(0).(*types.Event), args.Error(1)
}

// ListEvents for use with mock lib
func (c *MockClient) ListEvents(org string) ([]types.Event, error) {
	args := c.Called(org)
	return args.Get(0).([]types.Event), args.Error(1)
}

// DeleteEvent for use with mock lib
func (c *MockClient) DeleteEvent(entity, check string) error {
	args := c.Called(entity, check)
	return args.Error(0)
}

// ResolveEvent for use with mock lib
func (c *MockClient) ResolveEvent(event *types.Event) error {
	args := c.Called(event)
	return args.Error(0)
}
