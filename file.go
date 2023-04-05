/*
Copyright © 2023 Patrick Hermann patrick.hermann@sva.de
*/

package base

import (
	"fmt"
	"os"
)

func WriteDataToFile(outputFilePath, outputData string) bool {

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
