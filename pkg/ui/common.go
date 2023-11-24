package ui

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"

	"github.com/hawkv6/dummy-controller/internal/config"
	"github.com/manifoldco/promptui"
)

const (
	ReorderSids     = "Reorder SIDs"
	ChangeSidValues = "Change SID values"
	AddNewSid       = "Add new SID"
	DeleteSid       = "Delete SID"
)

const (
	AddFront = "Add to front"
	AddBack  = "Add to back"
)

func getAddingPositionList() []string {
	return []string{AddFront, AddBack}
}

func promptSelectAddingPosition() promptui.Select {
	return promptui.Select{
		Label: "Where would you like to add the SID?",
		Items: getAddingPositionList(),
	}
}

func getActionList() []string {
	return []string{ReorderSids, ChangeSidValues, AddNewSid, DeleteSid}
}

func promptSelectAction() promptui.Select {
	return promptui.Select{
		Label: "What would you like to do?",
		Items: getActionList(),
	}
}

func getServiceNames() []string {
	serviceNames := []string{}
	for k := range config.Params.Services {
		serviceNames = append(serviceNames, k)
	}
	return serviceNames
}

func promptSelectService() promptui.Select {
	return promptui.Select{
		Label: "Select a service to edit",
		Items: getServiceNames(),
	}
}

func getSidList(domainName string, intentList []string) []string {
	var sidList []string
	for _, intent := range config.Params.Services[domainName].Intents {
		if reflect.DeepEqual(intent.IntentList, intentList) {
			sidList = append(sidList, intent.Sid...)
		}
	}
	return sidList
}

func promptSelectSidList(domainName string, intentList []string) promptui.Select {
	return promptui.Select{
		Label: "Select a SID",
		Items: getSidList(domainName, intentList),
	}
}

func getIntentList(domainName string) []string {
	var intentList []string
	for _, intent := range config.Params.Services[domainName].Intents {
		joinedIntentList := strings.Join(intent.IntentList, ", ")
		intentList = append(intentList, joinedIntentList)
	}
	return intentList
}

func promptSelectIntentList(domainName string) promptui.Select {
	return promptui.Select{
		Label: "Select an Intent List to edit",
		Items: getIntentList(domainName),
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func prettyPrintService(domainName string) {
	for _, intent := range config.Params.Services[domainName].Intents {
		fmt.Printf("Intent List: %s\n", getIntentList(domainName))
		fmt.Printf("SIDs: %v\n", intent.Sid)
	}
}
