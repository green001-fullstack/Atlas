package domain

import "github.com/green001-fullstack/atlas/backend/internal/applicant/domain/valueobjects"

type Repository interface {
    Save(applicant *Applicant) error
    FindByID(id string) (*Applicant, error)
	ExistsByEmail( email valueobjects.EmailAddress)(bool, error)
}