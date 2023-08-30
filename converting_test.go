/*
Copyright © 2023 Patrick Hermann patrick.hermann@sva.de
*/

package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyIfStringIsBase64(t *testing.T) {

	assert := assert.New(t)

	isBase64 := VerifyIfStringIsBase64("aGVsbG8K")
	isNotBase64 := VerifyIfStringIsBase64("hello")

	assert.Equal(isBase64, true)
	assert.Equal(isNotBase64, false)

}

func TestDecodeBase64String(t *testing.T) {

	assert := assert.New(t)

	decodedBase64String := DecodeBase64String("aGVsbG8K")

	assert.Equal(decodedBase64String, "hello")

}
