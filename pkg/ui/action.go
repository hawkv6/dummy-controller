package ui

import (
	"github.com/hawkv6/dummy-controller/internal/config"
)

func deleteSid(domainName string, intent string, sid string) {
	for i, service := range config.Params.Services[domainName] {
		if service.Intent == intent {
			for j, s := range service.Sid {
				if s == sid {
					config.Params.Services[domainName][i].Sid = append(service.Sid[:j], service.Sid[j+1:]...)
					return
				}
			}
		}
	}
}

func changeSidValue(domainName string, intent string, sid string, newValue string) {
	for i, service := range config.Params.Services[domainName] {
		if service.Intent == intent {
			for j, s := range service.Sid {
				if s == sid {
					config.Params.Services[domainName][i].Sid[j] = newValue
					return
				}
			}
		}
	}
}

func reorderSids(domainName string, intent string) {
	for i, service := range config.Params.Services[domainName] {
		if service.Intent == intent {
			config.Params.Services[domainName][i].Sid = append(service.Sid[1:], service.Sid[0])
			return
		}
	}
}

func addToPosition(domainName string, intent string, sid string, position string) {
	newSidList := []string{}
	for i, service := range config.Params.Services[domainName] {
		if service.Intent == intent {
			switch position {
			case AddFront:
				newSidList = append(newSidList, sid)
				newSidList = append(newSidList, service.Sid...)
			case AddBack:
				newSidList = append(newSidList, service.Sid...)
				newSidList = append(newSidList, sid)
			}
			config.Params.Services[domainName][i].Sid = newSidList
		}
	}
}
