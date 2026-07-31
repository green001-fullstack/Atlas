package valueobjects

import (
	"net/mail"
	"strings"
)

func validate(text string) error{

	if text == "" {
		return ErrInvalidEmailAddress
	}
	addr, err := mail.ParseAddress(text)
	if err != nil{
		return ErrInvalidEmailAddress
	}
	if addr.Address != text{
		return ErrInvalidEmailAddress
	}
	
	return nil
}

func normalize(text string) string{
	value := strings.TrimSpace(text)
	return strings.ToLower(value)
}

type EmailAddress struct{
	value string
}

func NewEmailAddress(value string) (EmailAddress, error){
	normalizedEmail := normalize(value)
	if err := validate(normalizedEmail); err != nil{
		return EmailAddress{}, ErrInvalidEmailAddress
	}
	
	return EmailAddress{
		value: normalizedEmail,
	}, nil
}

func(e EmailAddress) String() string{
	return e.value
}