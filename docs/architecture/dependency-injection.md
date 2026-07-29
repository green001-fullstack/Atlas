# Dependency Injection

## Purpose

Dependency Injection supplies an object with the dependencies it requires instead of allowing it to create them itself.

This keeps the Application and Domain layers independent of infrastructure technologies.

---

## Benefits

- Loose coupling
- Easier testing
- Easier replacement of infrastructure
- Explicit dependencies
- Supports Hexagonal Architecture

---

## Atlas Example

GenerateTrustAssessmentService depends on:

- ApplicantRepository
- EvidenceRepository
- AssessmentRepository
- NotificationSender

These dependencies are created in `main.go` and injected into the service.