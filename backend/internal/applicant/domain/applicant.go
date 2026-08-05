package domain

import(
	"github.com/green001-fullstack/atlas/backend/internal/applicant/domain/valueobjects"
	"time"
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

	applicant.record(ApplicantRegistered{
	applicantID: applicant.id,
	email:	applicant.Email().String(),
	name: applicant.Name().String(),
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

// A function that helps to add domain events
func (a *Applicant) record(event DomainEvent) {
	a.domainEvents = append(a.domainEvents, event)
}

func (a *Applicant) GrantConsent(version string, grantedAt time.Time) error{
	if a.HasConsent(){
		return ErrConsentAlreadyGranted
	}

	consent, err := NewConsent(version, grantedAt)
	if err != nil{
		return err
	}

	a.consent = consent

	a.record(ConsentGranted{
		applicantID : a.id,
		version : version, 
		grantedAt: grantedAt,
	}) 
	return nil
}

func (a *Applicant) WithdrawConsent( withdrawnAt time.Time) error{
	if a.consent == nil {
		return ErrConsentNotGranted
	}

	if err := a.consent.Withdraw(withdrawnAt); err != nil {
		return err
	}

	a.record(ConsentWithdrawn{
		applicantID: a.id,
		withdrawnAt: withdrawnAt,
	})
	return nil
}