# Transactions

## Purpose

A transaction ensures that a business operation either completes entirely or leaves the system unchanged.

---

## Atlas Examples

### Generate Trust Assessment

Within one transaction:

- Save Trust Assessment
- Update Applicant Profile
- Persist Domain Events (for later publication)

If any step fails:

- Rollback all changes.

---

## Unit of Work

Responsibilities:

- Track modified Aggregates.
- Track newly created Aggregates.
- Commit all changes together.
- Roll back on failure.