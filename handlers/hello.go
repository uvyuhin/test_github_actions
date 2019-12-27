package handlers

import "log"

func SomeNumber(input int) int {
	result := input + 42
	log.Printf("result: %v", result)
	return result
}
