/*
Copyright © 2023 Patrick Hermann patrick.hermann@sva.de
*/

package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyIfFileOrDirExists(t *testing.T) {

	assert := assert.New(t)

	dirExists, _ := VerifyIfFileOrDirExists("/tmp", "dir")
	assert.Equal(dirExists, true)

	fileExists, _ := VerifyIfFileOrDirExists("/etc/os-release", "file")
	assert.Equal(fileExists, true)

}

func TestVerifyFileExistence(t *testing.T) {

	assert := assert.New(t)

	log := StdOutFileLogger("/dev/null", "2006-01-02 15:04:05", 50, 3, 28)

	fileExist := VerifyFileExistence("/tmp", log, true)

	assert.Equal(fileExist, true)

}

func TestWriteDataToFile(t *testing.T) {

	assert := assert.New(t)

	WriteDataToFile("/tmp/output-test", "hello")

	writtenText := ReadFileToVariable("/tmp/output-test")

	assert.Equal(writtenText, "hello")

}

func TestReadFileToVariable(t *testing.T) {

	assert := assert.New(t)

	WriteDataToFile("/tmp/output-test", "hello")

	writtenText := ReadFileToVariable("/tmp/output-test")

	assert.Equal(writtenText, "hello")
}

func TestCreateNestedDirectoryStructure(t *testing.T) {

	assert := assert.New(t)

	randomFolderName := GenerateRandomString(6)

	CreateNestedDirectoryStructure("/tmp/"+randomFolderName, 0600)

	folderExists, _ := VerifyIfFileOrDirExists("/tmp/"+randomFolderName, "dir")
	assert.Equal(true, folderExists)

}

func TestRemoveNestedFolder(t *testing.T) {

	var nestedFoldersRemoved bool

	assert := assert.New(t)

	randomFolderName := GenerateRandomString(6)

	CreateNestedDirectoryStructure("/tmp/"+randomFolderName, 0600)

	folderExists, _ := VerifyIfFileOrDirExists("/tmp/"+randomFolderName, "dir")

	if folderExists {
		RemoveNestedFolder("/tmp/" + randomFolderName)
		nestedFoldersRemoved = true
	}

	assert.Equal(true, nestedFoldersRemoved)

}
