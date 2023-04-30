/*
Copyright © 2023 Patrick Hermann patrick.hermann@sva.de
*/

package base

import "strconv"

func ConvertIntegerToString(inputNumber int) (outputNumber string) {

	outputNumber = strconv.FormatInt(int64(inputNumber), 10)

	return
}
