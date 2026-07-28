# Repository Design

Repositories provide access to Aggregate Roots.

Repositories do not contain business logic.

---

# ApplicantRepository

## Aggregate

Applicant

## Responsibilities

- Find Applicant by ID
- Save Applicant
- Find Applicant by Email

---

# AssessmentRepository

## Aggregate

Trust Assessment

## Responsibilities

- Find Assessment by ID
- Save Assessment
- Find Active Assessment

---

# FinancialEvidenceRepository

## Aggregate

Financial Evidence

## Responsibilities

- Save Evidence
- Find Evidence by Applicant
- Update Validation Status