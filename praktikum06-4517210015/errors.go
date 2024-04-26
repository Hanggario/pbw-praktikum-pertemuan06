package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func GetByID(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "GarinHanggario" {
		return NotFoundError
	}

	return nil
}

func main() {
	err := GetByID("GarinHanggario")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("not found error")
		} else {
			fmt.Println("unknown error")
		}
	} else {
		fmt.Println("No error")
	}
}
