package ui

// import (
// 	"fmt"

// 	"github.com/hawkv6/dummy-controller/pkg/api"
// 	"github.com/hawkv6/dummy-controller/pkg/intent"
// 	"github.com/hawkv6/dummy-controller/pkg/messaging"
// 	"github.com/manifoldco/promptui"
// )

// type UI struct {
// 	messagingChannels *messaging.MessagingChannels
// }

// func NewUI(messagingChannels *messaging.MessagingChannels) *UI {
// 	return &UI{
// 		messagingChannels: messagingChannels,
// 	}
// }

// func (ui *UI) Start() {
// 	for {
// 		prompt := promptSelectService()
// 		_, serviceName, err := prompt.Run()
// 		if err != nil {
// 			fmt.Printf("Prompt failed %v\n", err)
// 			return
// 		}

// 		fmt.Println("This is the current state of the service:")
// 		prettyPrintService(serviceName)

// 		prompt = promptSelectAction()
// 		_, action, err := prompt.Run()
// 		if err != nil {
// 			fmt.Printf("Prompt failed %v\n", err)
// 			return
// 		}
// 		prompt = promptSelectIntent(serviceName)
// 		_, intentName, err := prompt.Run()
// 		if err != nil {
// 			fmt.Printf("Prompt failed %v\n", err)
// 			return
// 		}

// 		switch action {
// 		case ReorderSids:
// 			reorderSids(serviceName, intentName)

// 		case ChangeSidValues:
// 			prompt = promptSelectSidList(serviceName, intentName)
// 			_, sid, err := prompt.Run()
// 			if err != nil {
// 				fmt.Printf("Prompt failed %v\n", err)
// 				return
// 			}

// 			newSidPrompt := promptui.Prompt{
// 				Label: "Enter new value for SID",
// 			}
// 			newValue, err := newSidPrompt.Run()
// 			if err != nil {
// 				fmt.Printf("Prompt failed %v\n", err)
// 				return
// 			}

// 			changeSidValue(serviceName, intentName, sid, newValue)

// 		case AddNewSid:
// 			prompt = promptSelectAddingPosition()
// 			_, position, err := prompt.Run()
// 			if err != nil {
// 				fmt.Printf("Prompt failed %v\n", err)
// 				return
// 			}

// 			newSidPrompt := promptui.Prompt{
// 				Label: "Enter new SID",
// 			}
// 			newSid, err := newSidPrompt.Run()
// 			if err != nil {
// 				fmt.Printf("Prompt failed %v\n", err)
// 				return
// 			}
// 			addToPosition(serviceName, intentName, newSid, position)

// 		case DeleteSid:
// 			prompt = promptSelectSidList(serviceName, intentName)
// 			_, sid, err := prompt.Run()
// 			if err != nil {
// 				fmt.Printf("Prompt failed %v\n", err)
// 				return
// 			}
// 			deleteSid(serviceName, intentName, sid)
// 		}

// 		// clearScreen()
// 		fmt.Println("This is the new state of the service:")
// 		prettyPrintService(serviceName)

// 		sidList := intent.GetIntentSidList(serviceName, intentName)
// 		ui.messagingChannels.ChMessageIntentResponse <- &api.PathResult{
// 			Ipv6DestinationAddress: serviceName,
// 			Intents: []*api.Intent{
// 				{Type: intent.StringToIntentType(intentName)},
// 			},
// 			Ipv6SidAddresses: sidList,
// 		}
// 		fmt.Println("-------------------------------------")
// 	}
// }
