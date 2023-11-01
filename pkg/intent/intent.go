package intent

import (
	"fmt"

	"github.com/hawkv6/dummy-controller/internal/config"
)

func GetIntentDetails(domainName string, intentType string) ([]string, error) {
	for _, service := range config.Params.Services[domainName] {
		if service.Intent == intentType {
			return service.Sid, nil
		}
	}
	return nil, fmt.Errorf("no intent found for domain %s and intent type %s", domainName, intentType)
}
