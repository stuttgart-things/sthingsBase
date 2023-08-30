/*
Copyright © 2023 Patrick Hermann patrick.hermann@sva.de
*/

package base

import (
	"encoding/base64"
	"log"
	"strconv"
	"strings"
)

func ConvertIntegerToString(inputNumber int) (outputNumber string) {

	outputNumber = strconv.FormatInt(int64(inputNumber), 10)

	return
}

func ConvertStringToInteger(inputNumber string) (outputNumber int) {

	outputNumber, err := strconv.Atoi(inputNumber)

	if err != nil {
		return
	}

	return
}

func ConvertStringToBoolean(trueOrFalse string) (boolValue bool) {

	boolValue, err := strconv.ParseBool(trueOrFalse)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func VerifyIfStringIsBase64(s string) bool {
	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}

func DecodeBase64String(base64Input string) string {
	stringOutput, err := base64.StdEncoding.DecodeString(base64Input)
	if err != nil {
		panic(err)
	}

	return strings.Trim(string(stringOutput), "\n")
}
