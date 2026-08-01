package valueobjects

import "strings"

func normalizeName(text string) string {
	newText := strings.TrimSpace(text)
	newTextSlice := strings.Fields(newText)
	finalText := strings.Join(newTextSlice, " ")
	return finalText
}

func validateName(text string) error{
	if len(text) > 100 || len(text) == 0{
		return ErrInvalidName
	} else if strings.ContainsAny(text, "0123456789"){
		return ErrInvalidName
	} else{
		return nil
	}
}

type Name struct {
	value string
}

func NewName(name string) (Name, error) {
	normalizedName := normalizeName(name)
	err := validateName(normalizedName)
	if err != nil{
		return Name{}, err
	}
	return Name{
		value: normalizedName,
	}, nil
}

func(n Name) String() string{
	return n.value
}