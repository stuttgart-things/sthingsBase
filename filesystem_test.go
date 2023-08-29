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
