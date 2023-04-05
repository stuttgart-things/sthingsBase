/*
Copyright © 2023 Patrick Hermann patrick.hermann@sva.de
*/

package base

import (
	"regexp"
	"strings"
)

func GetRegexSubMatch(scanText, regexPattern string) (string, bool) {

	rgx := regexp.MustCompile(regexPattern)
	regexSubMatch := rgx.FindStringSubmatch(scanText)

	if len(regexSubMatch) == 0 {
		return "", false
	}

	return strings.Trim(regexSubMatch[1], " "), true
}
