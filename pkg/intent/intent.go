package intent

import (
	"slices"

	"github.com/hawkv6/dummy-controller/internal/config"
	"github.com/hawkv6/dummy-controller/pkg/api"
	"github.com/hawkv6/dummy-controller/pkg/messaging"
)

type IntentHandler struct {
	messagingChannels *messaging.MessagingChannels
}

func NewIntentHandler(messagingChannels *messaging.MessagingChannels) *IntentHandler {
	return &IntentHandler{
		messagingChannels: messagingChannels,
	}
}

func (i *IntentHandler) Start() {
	go i.handleRequest()
}

func (i *IntentHandler) handleRequest() {
	for {
		intentRequest := <-i.messagingChannels.ChMessageIntentRequest
		sidList := GetIntentSidList(intentRequest.Ipv6DestinationAddress, IntentTypeToString(intentRequest.Intents[0].Type))
		response := &api.PathResult{
			Ipv6DestinationAddress: intentRequest.Ipv6DestinationAddress,
			Intents: []*api.Intent{
				{Type: intentRequest.Intents[0].Type},
			},
			Ipv6SidAddresses: sidList,
		}
		i.messagingChannels.ChMessageIntentResponse <- response
	}
}

// func GetIntentDetails(domainName string, intentType string) ([]string, error) {
// 	for _, service := range config.Params.Services[domainName] {
// 		if service.Intent == intentType {
// 			return service.Sid, nil
// 		}
// 	}
// 	return nil, fmt.Errorf("no intent found for domain %s and intent type %s", domainName, intentType)
// }

func GetIntentSidList(destAddr string, intentType string) []string {
	for _, service := range config.Params.Services {
		if slices.Contains(service.Ipv6Addresses, destAddr) {
			for _, intent := range service.Intents {
				if intent.Intent == intentType {
					return intent.Sid
				}
			}
		}
	}
	return nil
}
