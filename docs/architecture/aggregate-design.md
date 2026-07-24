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

# Assessment Request (Root)

## Contains
- Loan Amount
- Purpose
- Requested Product
- Requested Date
- Status
- Applicant Reference
- Organization Reference


# Assessment Aggregate

## Aggregate Root
Assessment

## Contains
- Financial Trust Score
- Confidence Score
- Recommendation
- Explanation
- Fraud Alert

## Business Invariant
- There can only be one active assessment per assessment request
- Recommendation must always match the current Trust Score.

# Evidence Aggregate

## Aggregate Root
Financial Evidence

## Contains
- Document MetaData
- Validation Status
- DataSource

## Business Invariant
- Every Financial Evidence must belong to exactly one Applicant.
- Unvalidated evidence cannot participate in a Trust Assessment.

# Organization Aggregate

## Aggregate Root
Organization

## Contains
- OrganizationUsers
- Roles
- Permissions