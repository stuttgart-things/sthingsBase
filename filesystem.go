/*
Copyright © 2023 Patrick Hermann patrick.hermann@sva.de
*/

package base

import (
	"fmt"
	"log"
	"os"
)

func CreateNestedDirectoryStructure(directoryStructure string, permission int) {
	osModePermissions := os.FileMode(int(permission))

	if err := os.MkdirAll(directoryStructure, osModePermissions); err != nil {
		fmt.Println(err)
	}

}

func RemoveNestedFolder(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		log.Fatal(err)
	}
}

func VerifyFileExistence(filePath string, log *Logger, exitOnError bool) bool {

	exists := false

	fileExists, _ := VerifyIfFileOrDirExists(filePath, "file")

	if !fileExists {
		log.Warn(filePath + " was not found")

		if exitOnError {
			log.Error("exiting.. goodbye galaxy!")
			os.Exit(3)
		}

	} else {
		log.Info(filePath + " exists")
		exists = true
	}

	return exists

}

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

func DeleteFile(filePath string) bool {

	err := os.Remove(filePath)

	if err != nil {
		return false
	} else {
		return true
	}

}

func WriteDataToFile(outputFilePath, outputData string) bool {

	f, err := os.Create(outputFilePath)

	if err != nil {
		log.Fatal(err)
		return false
	}

	defer f.Close()

	_, err2 := f.WriteString(outputData)

	if err2 != nil {
		log.Fatal(err2)
		return false
	}

	return true
}

func MoveRenameFileOnFS(oldLocation, newLocation string) {

	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
}
