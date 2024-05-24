/*
Copyright © 2024 Patrick Hermann patrick.hermann@sva.de
*/

package base

import (
	"slices"
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

func TestUniqueSlice(t *testing.T) {

	assert := assert.New(t)

	wantStringSlice := []string{"abc", "cde", "efg"}
	wantIntSlice := []int{1, 2, 3, 4}

	resultStringSlice := UniqueSlice([]string{"abc", "cde", "efg", "efg", "abc", "cde"})
	resultIntSlice := UniqueSlice([]int{1, 1, 2, 2, 3, 3, 4})

	assert.Equal(wantStringSlice, resultStringSlice)
	assert.Equal(wantIntSlice, resultIntSlice)

}

func TestGetStringPointerValue(t *testing.T) {
	assert := assert.New(t)

	// Test when the input is a non-nil pointer
	str := "test string"
	strPtr := &str
	result := GetStringPointerValue(strPtr)
	assert.Equal(str, result, "GetStringPointerValue did not return the correct value for a non-nil pointer")

	// Test when the input is a nil pointer
	var nilPtr *string
	result = GetStringPointerValue(nilPtr)
	assert.Equal("", result, "GetStringPointerValue did not return the correct value for a nil pointer")
}

func TestConvertIntegerToString(t *testing.T) {
	assert := assert.New(t)

	result := ConvertIntegerToString(123)
	assert.Equal("123", result, "ConvertIntegerToString did not convert correctly")
}

func TestConvertStringToInteger(t *testing.T) {
	assert := assert.New(t)

	result := ConvertStringToInteger("123")
	assert.Equal(123, result, "ConvertStringToInteger did not convert correctly")
}

func TestConvertStringToBoolean(t *testing.T) {
	assert := assert.New(t)

	result := ConvertStringToBoolean("true")
	assert.Equal(true, result, "ConvertStringToBoolean did not convert correctly")
}

func TestGenerateRandomString(t *testing.T) {
	assert := assert.New(t)

	result := GenerateRandomString(10)
	assert.Equal(10, len(result), "GenerateRandomString did not generate string of correct length")
}

func TestConvertStringToPointer(t *testing.T) {
	assert := assert.New(t)

	str := "test string"
	result := ConvertStringToPointer(str)
	assert.Equal(&str, result, "ConvertStringToPointer did not return the correct pointer")
}

func TestGetRandomPickFromSlice(t *testing.T) {
	// Define a slice of strings
	slice := []string{"Hello", "World", "from", "Go"}

	// Call the function
	result := GetRandomPickFromSlice(slice)

	// Check the result
	if !slices.Contains(slice, result) {
		t.Errorf("Expected 'World', got '%s'", result)
	}
}
