package valueobjects

import (
	"regexp"
	"strings"
)

var digitsOnly = regexp.MustCompile(`[^0-9]`)

func normalizePhoneNumber(text string) (string, error) {

	finalNumber := digitsOnly.ReplaceAllString(text, "")

	var transformedNumber string
	if len(finalNumber) == 11 {
		transformedNumber = "+234" + finalNumber[1:]
		return transformedNumber, nil
	}

	if len(finalNumber) == 13 {
		transformedNumber = "+" + finalNumber
		return transformedNumber, nil
	}

	return "", ErrInvalidPhoneNumber
}

func validatePhoneNumber(text string) error {

	if len(text) != 14 {
		return ErrInvalidPhoneNumber
	}

	validPrefixes := []string{
		"+23470",
		"+23480",
		"+23481",
		"+23490",
		"+23491",
	}

		for _, prefix := range validPrefixes {
			if strings.HasPrefix(text, prefix) {
				return nil
			}
		}

	return ErrInvalidPhoneNumber
}

type PhoneNumber struct {
	value string
}

func NewPhoneNumber(value string) (PhoneNumber, error) {
	normalizedPhoneNumber, err := normalizePhoneNumber(value)
	if err != nil {
		return PhoneNumber{}, err
	}
	err = validatePhoneNumber(normalizedPhoneNumber)
	if err != nil {
		return PhoneNumber{}, err
	}
	return PhoneNumber{value: normalizedPhoneNumber}, nil
}

func (p PhoneNumber) String() string {
	return p.value
}
