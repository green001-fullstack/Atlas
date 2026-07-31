package valueobjects

import "strings"

func normalizePhoneNumber(text string) string{
	newText := strings.ReplaceAll(text, " ", "")
	finalText := strings.ReplaceAll(newText, "-", "")
	return finalText
}

func validatePhoneNumber(text string) error{
	return nil
}

type PhoneNumber struct{
	value string
}

func NewPhoneNumber( value string) (PhoneNumber, error){
	return PhoneNumber{ value: value}, nil
}

func(p PhoneNumber) String() string{
	return p.value
}