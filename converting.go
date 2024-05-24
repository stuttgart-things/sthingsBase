/*
Copyright © 2024 Patrick Hermann patrick.hermann@sva.de
*/

package base

import (
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
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

func GenerateRandomString(length int) string {

	rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length+2)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[2 : length+2]
}

func MergeMaps[K comparable, V any](m1 map[K]V, m2 map[K]V) map[K]V {
	merged := make(map[K]V)
	for key, value := range m1 {
		merged[key] = value
	}
	for key, value := range m2 {
		merged[key] = value
	}
	return merged
}

// REMOVE DUPLICATES FROM ANY SLICE USING GENERICS
func UniqueSlice[T comparable](s []T) []T {
	inResult := make(map[T]bool)
	var result []T
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

// GET VALUES FROM STRING POINTER
func GetStringPointerValue(s *string) string {
	if s != nil {
		return *s
	}

	return ""
}

// CONVERT STRING TO STRING POINTER
func ConvertStringToPointer(s string) *string {
	return &s
}

// GET RANDOM ELEMENT FROM SLICE
func GetRandomPickFromSlice(slice []string) string {
	// GET RANDOM ELEMENT
	randomIndex := rand.Intn(len(slice))
	return slice[randomIndex]
}
