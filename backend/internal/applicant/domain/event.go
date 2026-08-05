package domain

import "time"

type DomainEvent interface{
	EventType() string
}

type ApplicantRegistered struct{
	applicantID string
	email       string
	name        string
}

func (e ApplicantRegistered) EventType() string {
	return "ApplicantRegistered"
}

type ConsentGranted struct{
	applicantID string
    version     string
	grantedAt	time.Time
}

func (c ConsentGranted) EventType() string {
	return "ConsentGranted"
}

type ConsentWithdrawn struct {
	applicantID string
	withdrawnAt time.Time
}

func (e ConsentWithdrawn) EventType() string {
	return "ConsentWithdrawn"
}