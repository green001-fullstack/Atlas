# Applicant Aggregate

## Aggregate Root
Applicant

## Contains
- EmailAddress
- PhoneNumber
- Consent
- FinancialProfile

## Business Invariants
- Applicant must grant consent before assessment.
- Applicant must always have at least one contact method.

## External References
- Assessment Requests
- Financial Evidence