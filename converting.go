/*
Copyright © 2023 Patrick Hermann patrick.hermann@sva.de
*/

package base

import "strconv"

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
