/*
Copyright © 2023 Patrick Hermann patrick.hermann@sva.de
*/

package base

import (
	"fmt"
	"log"
	"os"
)

func StoreVariableInFile(outputFilePath, outputData string) bool {

	f, err := os.Create(outputFilePath)

	if err != nil {
		fmt.Println(err)
		return false
	}

	defer f.Close()

	_, err2 := f.WriteString(outputData)

	if err2 != nil {
		fmt.Println(err2)
		return false
	}

	return true
}

func ReadFileToVariable(filePath string) string {

	content, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

/*
checks if a file exists or not
use w/ sthingsBase.FileExists("/home/my-file.yaml")
*/
func VerifyIfFileOrDirExists(filePath, kind string) (bool, error) {
	info, err := os.Stat(filePath)

	// returns true if folder exists
	if err == nil && kind == "dir" {
		return info.IsDir(), nil
	}

	// returns true if file exists
	if info != nil && kind == "file" {
		return true, nil
	}

	return false, err
}
