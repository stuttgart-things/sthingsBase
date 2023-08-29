/*
Copyright © 2023 Patrick Hermann patrick.hermann@sva.de
*/

package base

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyIfFileOrDirExists(t *testing.T) {
	assert := assert.New(t)

	DirExists, _ := VerifyIfFileOrDirExists("/tmp", "dir")
	assert.Equal(DirExists, true)

	FileExists, _ := VerifyIfFileOrDirExists("/etc/os-release", "file")
	assert.Equal(FileExists, true)

}

func TestVerifyFileExistence(t *testing.T) {

	log := StdOutFileLogger("/dev/null", "2006-01-02 15:04:05", 50, 3, 28)

	fileExist := VerifyFileExistence("/tmp", log, true)

	fmt.Println(fileExist)

}
