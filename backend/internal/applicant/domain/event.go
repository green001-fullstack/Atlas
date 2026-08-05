package domain

type DomainEvent interface{
	EventType() string
}

type ApplicantRegistered struct{
	applicantID string
	email       string
	name        string
}

func (e ApplicantRegistered) EventType() string {
	return "Applicant Registered"
}
