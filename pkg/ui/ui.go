package ui

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func Test() {
	for {
		prompt := promptSelectService()
		_, serviceName, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Println("This is the current state of the service:")
		prettyPrintService(serviceName)

		prompt = promptSelectAction()
		_, action, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch action {
		case ReorderSids:
			prompt := promptSelectIntent(serviceName)
			_, intent, err := prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
			reorderSids(serviceName, intent)

		case ChangeSidValues:
			prompt := promptSelectIntent(serviceName)
			_, intent, err := prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			prompt = promptSelectSidList(serviceName, intent)
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

			changeSidValue(serviceName, intent, sid, newValue)

		case AddNewSid:
			prompt := promptSelectIntent(serviceName)
			_, intent, err := prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
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
			addToPosition(serviceName, intent, newSid, position)

		case DeleteSid:
			prompt := promptSelectIntent(serviceName)
			_, intent, err := prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
			prompt = promptSelectSidList(serviceName, intent)
			_, sid, err := prompt.Run()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
			deleteSid(serviceName, intent, sid)
		}

		// clearScreen()
		fmt.Println("This is the new state of the service:")
		prettyPrintService(serviceName)
	}
}
