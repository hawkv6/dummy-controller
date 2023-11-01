package ui

import (
	"fmt"
	"os"
	"os/exec"

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

func getSidList(domainName string, intentType string) []string {
	sidList := []string{}
	for _, sid := range config.Params.Services[domainName] {
		if sid.Intent == intentType {
			sidList = append(sidList, sid.Sid...)
		}
	}
	return sidList
}

func promptSelectSidList(domainName string, intentType string) promptui.Select {
	return promptui.Select{
		Label: "Select a SID",
		Items: getSidList(domainName, intentType),
	}
}

func getIntentList(domainName string) []string {
	intentList := []string{}
	for _, intent := range config.Params.Services[domainName] {
		intentList = append(intentList, intent.Intent)
	}
	return intentList
}

func promptSelectIntent(domainName string) promptui.Select {
	return promptui.Select{
		Label: "Select an Intent",
		Items: getIntentList(domainName),
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func prettyPrintService(domainName string) {
	for _, service := range config.Params.Services[domainName] {
		fmt.Printf("Intent: %s\n", service.Intent)
		fmt.Printf("SIDs: %v\n", service.Sid)
	}
}
