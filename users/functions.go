package users

import (
	"errors"
    "strings"
	"fmt"
)

func Length(str string) int {
	symbolCount := len([]rune(str))
	return symbolCount
}

func IsEmpty(str, val string) error {
	if strings.TrimSpace(str) == "" {
		return fmt.Errorf("%s should not be empty", val)
	}
	return nil
}

func LenName(name string) error {
	if Length(name) >= 4 && Length(name) <= 20 {
		return nil
	}
	return errors.New("name should be between 4 and 20 symbols long")
}

func LenPassword(password string) error {
	if Length(password) >= 8 && Length(password) <= 20 {
		return nil
	}
	return errors.New("name should be between 8 and 20 symbols long")
}

