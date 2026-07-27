# GenerateTrustAssessment

## Purpose
Requests the generation of a Trust Assessment for a submitted Assessment Request.

## Initiated By
Loan Officer

## Target Aggregate
Trust Assessment Request

## Expected Outcome
If successful, a Trust Assessment is generated and a TrustAssessmentGenerated domain event is published.

## Possible Failures
- Assessment Request not found.
- Assessment already in progress.
- Required evidence missing.
- Applicant consent not granted.


# Register Applicant

## Purpose
Registers a new applicant so they can request financial trust assessments.

## Initiated By
Applicant

## Target Aggregate
Applicant Aggregate

## Expected Outcome
- Applicant is created.
- Applicant profile is initialized.
- ApplicantRegistered domain event is published.

## Possible Failures
- Applicant already exists.
- Required fields missing.
- Invalid email address.
- Invalid phone number.

# GrantConsent

## Purpose
Applicant needs to give consent that the organization can make use of his/her information as well as uploaded document during assessment

## Initiated By
Applicant

## Target Aggregate
Applicant Aggregate

## Expected Outcome
- Consent is recorded.
- ConsentGranted domain event is published.

## Possible Failures
- Applicant not found.
- Consent already granted.
- Invalid consent request.

# RequestAssessment

## Purpose
Creates a new request for a financial trust assessment.

## Initiated By
Loan Officer

## Target Aggregate
Assessment Request Aggregate

## Expected Outcome
- Assessment Request is created.
- AssessmentRequested domain event is published.

## Possible Failures
- Applicant not found.
- Consent not granted.
- Active assessment already exists.

# UploadFinancialEvidence

## Purpose
Uploads financial evidence that will be used during trust assessment.

## Initiated By
Applicant

## Target Aggregate
Financial Evidence Aggregate

## Expected Outcome
- Evidence stored successfully.
- FinancialEvidenceUploaded domain event published.

## Possible Failures
- Unsupported document type.
- File too large.
- Upload failed.
- Applicant not found.

# ValidateFinancialEvidence

## Purpose
Verifies that uploaded financial evidence is authentic and usable.

## Initiated By
System

## Target Aggregate
Financial Evidence Aggregate

## Expected Outcome
- Evidence marked as validated.
- FinancialEvidenceValidated domain event published.

## Possible Failures
- Document unreadable.
- Evidence does not belong to applicant.
- Validation service unavailable.

# CompleteAssessment

## Purpose
Marks a trust assessment as complete.

## Initiated By
System

## Target Aggregate
Trust Assessment Aggregate

## Expected Outcome
- Assessment completed.
- AssessmentCompleted domain event published.

## Possible Failures
- Assessment not found.
- Assessment already completed.
- Recommendation missing.

# CancelAssessment

## Purpose
Cancels an active assessment request.

## Initiated By
Loan Officer

## Target Aggregate
Assessment Request Aggregate

## Expected Outcome
- Assessment Request cancelled.
- AssessmentCancelled domain event published.

## Possible Failures
- Assessment already completed.
- Assessment already cancelled.
- Assessment not found.

# ResolveFraudAlert

## Purpose
Marks a fraud investigation as resolved.

## Initiated By
Risk Analyst

## Target Aggregate
Fraud Alert Aggregate

## Expected Outcome
- Fraud Alert resolved.
- Audit log updated.

## Possible Failures
- Fraud Alert not found.
- Fraud Alert already resolved.

# InviteOrganizationUser

## Purpose
Invites a new user to join an organization.

## Initiated By
Organization Administrator

## Target Aggregate
Organization Aggregate

## Expected Outcome
- Invitation created.
- Organization user notified.

## Possible Failures
- User already exists.
- Organization not found.
- Invitation already pending.