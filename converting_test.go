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

func TestMergeMaps(t *testing.T) {

	// WANTED RESULT
	want := make(map[string]interface{})
	want["ram"] = 4096
	want["cpu"] = 4

	// DEFINE TEST VALUES
	defaultVariables := make(map[string]interface{})
	envVariables := make(map[string]interface{})
	defaultVariables["ram"] = 2048
	defaultVariables["cpu"] = 4
	envVariables["ram"] = 4096

	// MERGE DEFAULTS + VALUES
	got := MergeMaps(defaultVariables, envVariables)

	assert.InDeltaMapValues(t, got, want, 0.0, "Word count wrong. Got %v, want %v", got, want)
}
