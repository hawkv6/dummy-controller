package ui

import (
	"fmt"
	"strings"

	"github.com/hawkv6/dummy-controller/pkg/intent"
	"github.com/hawkv6/dummy-controller/pkg/messaging"
	"github.com/manifoldco/promptui"
)

type UI struct {
	messagingChannels *messaging.MessagingChannels
}

func NewUI(messagingChannels *messaging.MessagingChannels) *UI {
	return &UI{
		messagingChannels: messagingChannels,
	}
}

func (ui *UI) Start() {
	for {
		prompt := promptSelectService()
		_, serviceName, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		prompt = promptSelectAction()
		_, action, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		prompt = promptSelectIntentList(serviceName)
		_, intentListString, err := prompt.Run()
		intentList := strings.Split(intentListString, ", ")
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch action {
		case ReorderSids:
			reorderSids(serviceName, intentList)

		case ChangeSidValues:
			prompt = promptSelectSidList(serviceName, intentList)
			_, sid, err := prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			newSidPrompt := promptui.Prompt{
				Label: "Enter new value for SID",
			}
			newValue, err := newSidPrompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			changeSidValue(serviceName, intentList, sid, newValue)

		case AddNewSid:
			prompt = promptSelectAddingPosition()
			_, position, err := prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			newSidPrompt := promptui.Prompt{
				Label: "Enter new SID",
			}
			newSid, err := newSidPrompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
			addToPosition(serviceName, intentList, newSid, position)

		case DeleteSid:
			prompt = promptSelectSidList(serviceName, intentList)
			_, sid, err := prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
			deleteSid(serviceName, intentList, sid)
		}

		clearScreen()

		destinationAddresses := intent.GetIpv6Addresses(serviceName)
		pathResults := intent.CreatePathResults(destinationAddresses, intentList)

		for _, pathResult := range pathResults {
			ui.messagingChannels.ChMessageIntentResponse <- pathResult
		}
	}
}
