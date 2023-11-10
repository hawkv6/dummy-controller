package ui

import (
	"github.com/hawkv6/dummy-controller/internal/config"
)

func deleteSid(domainName string, intent string, sid string) {
	for k, i := range config.Params.Services[domainName].Intents {
		if i.Intent == intent {
			for j, s := range i.Sid {
				if s == sid {
					config.Params.Services[domainName].Intents[k].Sid = append(i.Sid[:j], i.Sid[j+1:]...)
					return
				}
			}
		}
	}
}

func changeSidValue(domainName string, intent string, sid string, newValue string) {
	for k, i := range config.Params.Services[domainName].Intents {
		if i.Intent == intent {
			for j, s := range i.Sid {
				if s == sid {
					config.Params.Services[domainName].Intents[k].Sid[j] = newValue
					return
				}
			}
		}
	}
}

func reorderSids(domainName string, intent string) {
	for k, i := range config.Params.Services[domainName].Intents {
		if i.Intent == intent {
			config.Params.Services[domainName].Intents[k].Sid = append(i.Sid[1:], i.Sid[0])
			return
		}
	}
}

func addToPosition(domainName string, intent string, sid string, position string) {
	newSidList := []string{}
	for k, i := range config.Params.Services[domainName].Intents {
		if i.Intent == intent {
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
