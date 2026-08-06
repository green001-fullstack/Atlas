package handlers

import (
	"github.com/green001-fullstack/atlas/backend/internal/applicant/domain"
	"github.com/green001-fullstack/atlas/backend/internal/applicant/application/ports"
	"github.com/green001-fullstack/atlas/backend/internal/applicant/application/commands"
)

type RegisterApplicantHandler struct {
    repository     domain.Repository
    eventPublisher ports.EventPublisher
}

type RegisterApplicantResult struct {
    ApplicantID string
}

func NewRegisterApplicantHandler( repository domain.Repository, eventPublisher ports.EventPublisher) *RegisterApplicantHandler {
    return &RegisterApplicantHandler{
        repository: repository,
        eventPublisher: eventPublisher,
    }
}

func (h *RegisterApplicantHandler) Handle( command commands.RegisterApplicantCommand ) (RegisterApplicantResult, error)
