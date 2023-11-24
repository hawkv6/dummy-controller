package intent

import (
	"reflect"
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
		response := CreatePathResultFromRequest(intentRequest)
		i.messagingChannels.ChMessageIntentResponse <- response
	}
}

// func GetIntentSidList(destAddr string, intentType string) []string {
// 	for _, service := range config.Params.Services {
// 		if slices.Contains(service.Ipv6Addresses, destAddr) {
// 			for _, intent := range service.Intents {
// 				if intent.Intent == intentType {
// 					return intent.Sid
// 				}
// 			}
// 		}
// 	}
// 	return nil
// }

// func GetIpv6Address(serviceName string) string {
// 	for key, service := range config.Params.Services {
// 		if key == serviceName {
// 			return service.Ipv6Addresses[0]
// 		}
// 	}
// 	return ""
// }

func GetIpv6Addresses(serviceName string) []string {
	for key, service := range config.Params.Services {
		if key == serviceName {
			return service.Ipv6Addresses
		}
	}
	return nil
}

func CreatePathResults(daddresses []string, intentList []string) []*api.PathResult {
	sidList := GetSidList(daddresses[0], intentList)
	var pathResults []*api.PathResult
	var intents []*api.Intent
	for _, intent := range intentList {
		intents = append(intents, &api.Intent{
			Type: StringToIntentType(intent),
		})
	}
	for _, daddr := range daddresses {
		pathResults = append(pathResults, &api.PathResult{
			Ipv6DestinationAddress: daddr,
			Intents:                intents,
			Ipv6SidAddresses:       sidList,
		})
	}
	return pathResults
}

func CreatePathResultFromRequest(request *api.PathRequest) *api.PathResult {
	sidList := GetSidList(request.Ipv6DestinationAddress, GetIntentsFromRequest(request))
	return &api.PathResult{
		Ipv6DestinationAddress: request.Ipv6DestinationAddress,
		Intents:                request.Intents,
		Ipv6SidAddresses:       sidList,
	}
}

func GetIntentsFromRequest(request *api.PathRequest) []string {
	var intentList []string
	for _, intent := range request.Intents {
		intentList = append(intentList, IntentTypeToString(intent.Type))
	}
	return intentList
}

func GetSidList(daddr string, intents []string) []string {
	var sidList []string

	for _, service := range config.Params.Services {
		if slices.Contains(service.Ipv6Addresses, daddr) {
			for _, intent := range service.Intents {
				if reflect.DeepEqual(intent.IntentList, intents) {
					sidList = append(sidList, intent.Sid...)
				}
			}
		}
	}

	return sidList
}
