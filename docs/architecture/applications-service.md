# Application Services

## Purpose

Application Services coordinate business use cases.

They:
- Receive Commands.
- Load Aggregates through Repositories.
- Invoke Domain behavior.
- Persist changes.
- Publish Domain Events.

They DO NOT:
- Contain business rules.
- Calculate Trust Scores.
- Make lending decisions.
- Access the database directly.

---

# RequestAssessmentService

## Input
RequestAssessmentCommand

## Responsibilities

1. Validate request format.
2. Load Applicant.
3. Delegate assessment request creation to the Domain.
4. Save changes.
5. Publish Domain Events.

---

# GenerateTrustAssessmentService

## Input
GenerateTrustAssessmentCommand

## Responsibilities

1. Load Assessment Request.
2. Load Applicant and Financial Evidence.
3. Invoke the Trust Assessment Aggregate.
4. Persist the updated Aggregate.
5. Publish TrustAssessmentGenerated.