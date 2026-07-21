# Business Concepts

Applicant

Loan Application

Organization

Financial Evidence

Financial Signal

Trust Assessment

Financial Trust Score

Recommendation

Explanation

Risk Indicator

Fraud Alert

Consent

Reviewer

Assessment Report

Financial Profile

Financial Behavior

Data Source

Assessment Request

Confidence Score

Financial Pattern



| Term | Meaning |
|------|---------|
| Applicant | An individual or business requesting a financial trust assessment. |
| Organization | A bank, fintech, cooperative, or lender using Atlas. |
| Financial Evidence | Raw financial information submitted or collected with consent. |
| Financial Behavior | Patterns inferred from financial evidence. |
| Trust Assessment | Atlas's evaluation of an applicant's financial behavior. |
| Recommendation | Guidance provided to a loan officer based on the assessment. |

                    Organization
                         │
                  has many Users
                         │
                         ▼
                Organization User
                         │
             requests Assessment
                         │
                         ▼
                Assessment Request
                    │          │
                    │          │
             for Applicant      │
                    │          │
                    ▼          │
                Applicant       │
                    │           │
        has many Financial Evidence
                    │
                    ▼
          Financial Evidence
                    │
             used to generate
                    ▼
            Trust Assessment
              │      │      │
              │      │      │
              ▼      ▼      ▼
      Recommendation
      Explanation
      Fraud Alert