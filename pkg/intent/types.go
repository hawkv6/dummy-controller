package intent

import "github.com/hawkv6/dummy-controller/pkg/api"

func IntentTypeToString(intentType api.IntentType) string {
	switch intentType {
	case api.IntentType_INTENT_TYPE_HIGH_BANDWIDTH:
		return "high-bandwidth"
	case api.IntentType_INTENT_TYPE_LOW_BANDWIDTH:
		return "low-bandwidth"
	case api.IntentType_INTENT_TYPE_LOW_LATENCY:
		return "low-latency"
	case api.IntentType_INTENT_TYPE_FLEX_ALGO:
		return "flex-algo"
	case api.IntentType_INTENT_TYPE_SFC:
		return "sfc"
	}

	return "unspecified"
}

func StringToIntentType(intentType string) api.IntentType {
	switch intentType {
	case "high-bandwidth":
		return api.IntentType_INTENT_TYPE_HIGH_BANDWIDTH
	case "low-bandwidth":
		return api.IntentType_INTENT_TYPE_LOW_BANDWIDTH
	case "low-latency":
		return api.IntentType_INTENT_TYPE_LOW_LATENCY
	case "flex-algo":
		return api.IntentType_INTENT_TYPE_FLEX_ALGO
	case "sfc":
		return api.IntentType_INTENT_TYPE_SFC
	}

	return api.IntentType_INTENT_TYPE_UNSPECIFIED
}
