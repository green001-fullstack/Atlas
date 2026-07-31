package valueobjects

type EmailAddress struct{
	value string
}

func NewEmailAddress(value string) (EmailAddress, error){
	return EmailAddress{
		value: value,
	}, nil
}