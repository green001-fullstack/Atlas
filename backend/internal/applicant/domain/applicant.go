package domain

import(
	"github.com/green001-fullstack/atlas/backend/internal/applicant/domain/valueobjects"
)

type Applicant struct{
	id string
	name valueobjects.Name
	email valueobjects.EmailAddress
	phoneNumber valueobjects.PhoneNumber
	consent *Consent
	domainEvents []DomainEvent
}

func NewApplicant( id string, name valueobjects.Name, email valueobjects.EmailAddress, phone valueobjects.PhoneNumber ) (*Applicant, error){
	if id == ""{
		return nil, ErrInvalidApplicant
	} 
	applicant := &Applicant{
		id : id,
		name : name,
		email : email,
		phoneNumber: phone,
	}

	applicant.domainEvents = append(applicant.domainEvents, ApplicantRegistered{
		applicantID: applicant.id,
	})

	return  applicant, nil
	}

func (a *Applicant) ID() string{ 
		return a.id 
	}

func(a *Applicant) Name() valueobjects.Name { 
	return a.name 
}

func(a *Applicant) Email() valueobjects.EmailAddress{ 
	return a.email
}

func(a *Applicant) PhoneNumber()valueobjects.PhoneNumber{
	return a.phoneNumber
}

func(a *Applicant) HasConsent() bool{
	return a.consent != nil && a.consent.IsActive()
}

func (a *Applicant) DomainEvents() []DomainEvent{
	// to prevent modifying the original slice
	events := make([]DomainEvent, len(a.domainEvents))
	copy(events, a.domainEvents)
	return events
	// instead of 
	// return a.domainEvents
}

func (a *Applicant) ClearDomainEvents() {
	a.domainEvents = nil
}
