/*
Copyright © 2024 Patrick Hermann patrick.hermann@sva.de
*/

package base

import (
	"slices"
)

func CheckForStringInSlice(iN []string, searchFor string) bool {

	if slices.Contains(iN, searchFor) {
		return true
	} else {
		return false
	}
}
