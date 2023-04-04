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
	rs := rgx.FindStringSubmatch(scanText)

	if len(rs) == 0 {
		return "", false
	}

	return strings.Trim(rs[1], " "), true
}
