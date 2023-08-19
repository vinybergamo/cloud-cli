package utils

import "errors"

func ValidateEmpty(input string, message string) error {
	if input == "" {
		return errors.New(message)
	}
	return nil
}
