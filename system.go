/*
Copyright © 2023 Patrick Hermann patrick.hermann@sva.de
*/

package base

import (
	"log"
	"os/exec"
	"os/user"
	"strings"

	"golang.org/x/sys/unix"
)

func ValidateOSUser() (username string) {

	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	username = user.Username

	return

}

func CheckForUnixWritePermissions(filePath string) bool {
	return unix.Access(filePath, unix.W_OK) == nil
}

func GetExternalProcessOutputToVar(cmd string, args []string) string {

	c, b := exec.Command(cmd, args...), new(strings.Builder)
	c.Stdout = b
	c.Run()
	print(b.String())

	return b.String()

}
