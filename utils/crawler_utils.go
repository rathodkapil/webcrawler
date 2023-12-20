package utils

import "strings"

func ToIgnore(link string) bool {
	byPassLinks := []string{"youtube", "twitter", "instagram", "facebook", "linkedin", "adjust"}
	for _, v := range byPassLinks {
		if strings.Contains(link, v) {
			return true
		}
	}
	return false
}
