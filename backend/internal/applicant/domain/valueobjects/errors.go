package valueobjects

import "errors"

var ErrInvalidEmailAddress = errors.New("invalid email address")
var ErrInvalidPhoneNumber = errors.New("invalid phone number")
var ErrInvalidName = errors.New("invalid name")