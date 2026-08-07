package handlers

import (
	"github.com/green001-fullstack/atlas/backend/internal/applicant/application/ports"
	"github.com/green001-fullstack/atlas/backend/internal/applicant/domain"
	"time"
	"github.com/green001-fullstack/atlas/backend/internal/applicant/application/commands"
	// "github.com/green001-fullstack/atlas/backend/internal/applicant/domain/valueobjects"
)
type GrantConsentHandler struct{
	repository domain.ApplicantRepository
	eventPublisher ports.EventPublisher
}

type GrantConsentResult struct{}

func NewGrantConsentHandler(repository domain.ApplicantRepository, eventPublisher ports.EventPublisher) *GrantConsentHandler {
	return &GrantConsentHandler{
		repository: repository,
		eventPublisher: eventPublisher,
	}
}


func (h *GrantConsentHandler) Handle(command commands.GrantConsentCommand) (GrantConsentResult, error){
	applicant, err := h.repository.FindByID(command.ApplicantID)

	if err != nil{
		return GrantConsentResult{}, domain.ErrApplicantNotFound
	}

	grantedAt := time.Now()

	err = applicant.GrantConsent(command.Version, grantedAt)
	if err := h.repository.Save(applicant); err != nil{
		return GrantConsentResult{}, err
	} 

	if err := h.eventPublisher.Publish(applicant.DomainEvents()); err != nil{
		return GrantConsentResult{}, err
	}
	applicant.ClearDomainEvents()

	return GrantConsentResult{}, nil
}