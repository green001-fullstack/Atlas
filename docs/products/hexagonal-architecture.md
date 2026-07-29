# Hexagonal Architecture

## Purpose

Hexagonal Architecture places the Domain at the center of the system.

External technologies communicate with the Domain through Ports (interfaces).

Adapters implement those interfaces.

---

## Inbound Adapters

- REST API
- React Frontend
- CLI
- Scheduled Jobs

---

## Outbound Ports

- Applicant Repository
- Assessment Repository
- Notification Sender
- Storage Service

---

## Outbound Adapters

- PostgreSQL Repository
- SMTP Email Sender
- AWS S3 Storage