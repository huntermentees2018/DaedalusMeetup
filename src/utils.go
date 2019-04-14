package src

import (
	"strings"
)

func RegexCommand(text string) string {
	if strings.Contains(text, "unsubscribe") {
		return "unsubscribe"
	} else if strings.Contains(text, "subscribe") {
		return "subscribe"
	} else if strings.Contains(text, "list") {
		return "list"
	} else if strings.Contains(text, "help") {
		return "help"
	}

	return "idk"
}
