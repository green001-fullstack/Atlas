# Applicant

## Description
Represents an individual or business requesting a financial trust assessment.

## Relationships
- Can have many Assessment Requests
- Can provide many Financial Evidence records
- Can have many Trust Assessments

## Business Rules
- Must grant consent before assessment.
- May apply through multiple organizations.
- Cannot have an assessment without sufficient evidence.
- An applicant's profile may evolve over time as new evidence becomes available.

# Organization

## Description
Represents a financial institution that uses Atlas to assess applicants.

## Relationships
- Has many Organization Users.
- Can submit many Assessment Requests.
- Can define lending policies (future feature).
- Can view many Assessment Reports.
 
# Organization users

## Description
This represents the various users of the Atlas

## Relationships
- Belongs to one Organization.
- Can submit many Assessment Requests.
- Can review many Trust Assessments.

# Business Rules
- Must belong to an Organization.
- Must have an assigned role.
- Permissions depend on role.

# Assessment Request

## Description
Represents when the applicant requests to be assessed for possible eligibility

## Relationships
- Assessment can be requested by applicants at different organisation at different times

## Business Rule
- An organization cannot create multiple active assessment requests for the same applicant.

# Financial Evidence

## Description
This represents the document of the applicant that will serve as a proof of their financial commitment

## Relationships
- The submitted evidence bears the details and financial commitment of the applicant

### Business Rule
- Evidence must be linked to one applicant.
- Evidence must originate from an approved data source.
- Evidence must pass validation before it can be used in an assessment.

# Financial Profile

## Description
Represents the evolving financial characteristics inferred from an applicant's financial evidence.

# Trust Assessment

## Description
Represents Atlas's evaluation of an applicant's financial behavior based on available financial evidence.

## Relationships
- Belongs to one Assessment Request.
- Uses many Financial Evidence records.
- Produces one Recommendation.
- Produces one Explanation.
- May produce one or more Fraud Alerts.

# Recommendation

## Description
Represents Atlas's suggested course of action based on a Trust Assessment.

## Business Rules
- Must belong to a Trust Assessment.
- Must include a confidence level.
- Must never represent the final lending decision

# Fraud Alert

## Description
Represents suspicious financial behavior identified during assessment.

## Relationships
Belongs to one Trust Assessment

## Business Rules
- Does not automatically reject an applicant.
- Must include supporting evidence.
- May require manual investigation.

# Explanation

## Description
Represents the reasons and evidence supporting a Trust Assessment.

## Relationships
- Belongs to one Trust Assessment.

## Business Rules:
- Every Trust Assessment must have an Explanation.
- Every explanation must reference observable financial evidence.
- Explanations must be understandable by non-technical users.