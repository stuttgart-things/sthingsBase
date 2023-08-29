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

func VerifyFileExistence(filePath, kind string, log *Logger, exitOnError bool) bool {

	exists := false

	fileExists, _ := VerifyIfFileOrDirExists(filePath, kind)
	if !fileExists {
		log.Warn(kind + " " + filePath + " was not found")

		if exitOnError {
			log.Error("exiting.. goodbye galaxy!")
			os.Exit(3)
		}

	} else {
		log.Info(kind + " " + filePath + " exists")
		exists = true
	}

	return exists

}
