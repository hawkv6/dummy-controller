package messaging

import "github.com/hawkv6/dummy-controller/pkg/api"

type MessagingChannels struct {
	ChMessageIntentRequest  chan *api.PathRequest
	ChMessageIntentResponse chan *api.PathResult
}

func NewMessagingChannels() *MessagingChannels {
	return &MessagingChannels{
		ChMessageIntentRequest:  make(chan *api.PathRequest),
		ChMessageIntentResponse: make(chan *api.PathResult),
	}
}
