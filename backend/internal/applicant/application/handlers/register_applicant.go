package handlers

import (
	"github.com/green001-fullstack/atlas/backend/internal/applicant/domain"
	"github.com/green001-fullstack/atlas/backend/internal/applicant/application/ports"
	"github.com/green001-fullstack/atlas/backend/internal/applicant/application/commands"
    "github.com/green001-fullstack/atlas/backend/internal/applicant/domain/valueobjects"
)

type RegisterApplicantHandler struct {
    repository     domain.ApplicantRepository
    eventPublisher ports.EventPublisher
}

type RegisterApplicantResult struct {
    ApplicantID string
}

func NewRegisterApplicantHandler( repository domain.ApplicantRepository, eventPublisher ports.EventPublisher) *RegisterApplicantHandler {
    return &RegisterApplicantHandler{
        repository: repository,
        eventPublisher: eventPublisher,
    }
}

func (h *RegisterApplicantHandler) Handle( command commands.RegisterApplicantCommand ) (RegisterApplicantResult, error){
    name, err := valueobjects.NewName(command.Name)
	if err != nil {
		return RegisterApplicantResult{}, err
	}

	email, err := valueobjects.NewEmailAddress(command.Email)
	if err != nil {
		return RegisterApplicantResult{}, err
	}

	phone, err := valueobjects.NewPhoneNumber(command.PhoneNumber)
	if err != nil {
		return RegisterApplicantResult{}, err
	}

    exists, err := h.repository.ExistsByEmail(email)
    if err != nil {
        return RegisterApplicantResult{}, err
    }

    if exists {
        return RegisterApplicantResult{}, domain.ErrApplicantAlreadyExists
    }

	id := h.repository.NextIdentity()
	applicant, err := domain.NewApplicant(id, name, email, phone)
    if err != nil{
        return RegisterApplicantResult{}, err
    }

	err = h.repository.Save(applicant)
    if err != nil{
        return RegisterApplicantResult{}, err
    }
	if err := h.eventPublisher.Publish(applicant.DomainEvents()); err != nil{
        return RegisterApplicantResult{}, err
    }

	applicant.ClearDomainEvents()

	return RegisterApplicantResult{
		ApplicantID: applicant.ID(),
	}, nil
}
