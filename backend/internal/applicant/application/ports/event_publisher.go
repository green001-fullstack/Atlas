package ports

import "github.com/green001-fullstack/atlas/backend/internal/applicant/domain"

type EventPublisher interface {
    Publish(events []domain.DomainEvent) error
}