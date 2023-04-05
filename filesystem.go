/*
Copyright © 2023 Patrick Hermann patrick.hermann@sva.de
*/

package base

import (
	"fmt"
	"os"
)

func CreateNestedDirectoryStructure(directoryStructure string, permission int) {

	osModePermissions := os.FileMode(int(permission))

	if err := os.MkdirAll(directoryStructure, osModePermissions); err != nil {
		fmt.Println(err)
	}

}
