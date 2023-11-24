package ui

import (
	"reflect"

	"github.com/hawkv6/dummy-controller/internal/config"
)

func deleteSid(domainName string, intentList []string, sid string) {
	for k, i := range config.Params.Services[domainName].Intents {
		if reflect.DeepEqual(i.IntentList, intentList) {
			for j, s := range i.Sid {
				if s == sid {
					config.Params.Services[domainName].Intents[k].Sid = append(i.Sid[:j], i.Sid[j+1:]...)
					return
				}
			}
		}
	}
}

func changeSidValue(domainName string, intentList []string, sid string, newValue string) {
	for k, i := range config.Params.Services[domainName].Intents {
		if reflect.DeepEqual(i.IntentList, intentList) {
			for j, s := range i.Sid {
				if s == sid {
					config.Params.Services[domainName].Intents[k].Sid[j] = newValue
					return
				}
			}
		}
	}
}

func reorderSids(domainName string, intentList []string) {
	for k, i := range config.Params.Services[domainName].Intents {
		if reflect.DeepEqual(i.IntentList, intentList) {
			config.Params.Services[domainName].Intents[k].Sid = append(i.Sid[1:], i.Sid[0])
			return
		}
	}
}

func addToPosition(domainName string, intentList []string, sid string, position string) {
	newSidList := []string{}
	for k, i := range config.Params.Services[domainName].Intents {
		if reflect.DeepEqual(i.IntentList, intentList) {
			switch position {
			case AddFront:
				newSidList = append(newSidList, sid)
				newSidList = append(newSidList, i.Sid...)
			case AddBack:
				newSidList = append(newSidList, i.Sid...)
				newSidList = append(newSidList, sid)
			}
			config.Params.Services[domainName].Intents[k].Sid = newSidList
		}
	}
}
