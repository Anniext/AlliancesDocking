package mqtt

import "strings"

func getCoreID(topic string) string {
	tps := strings.Split(topic, "/")
	if len(tps) > 2 {
		return tps[1]
	}
	return ""
}
