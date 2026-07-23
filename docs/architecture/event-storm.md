# Domain Events & Workflow

## 1. Domain Events

| Event | Description |
|---|---|
| **Applicant Registered** | A new applicant has created an account in the system. |
| **Consent Granted** | The applicant has authorized use of their financial data. |
| **Applicant Requested Assessment** | The applicant has submitted a request for a trust assessment. |
| **Assessment Put On Hold** | An in-progress assessment has been paused (e.g., pending more evidence). |
| **Fraud Analysis Completed** | The fraud-detection process has finished evaluating the applicant's data. |
| **Fraud Alert Raised** | Suspicious activity has been detected and flagged for investigation. |
| **Trust Assessment Generated** | A trust assessment has been produced, combining score, explanation, and recommendation. |

---

## 2. Detailed Workflow

The full end-to-end process, from applicant registration to final decision:

```
 1. Applicant
        │
        ▼
 2. Registers
        │
        ▼
 3. Grants Consent
        │
        ▼
 4. Uploads Financial Evidence
        │
        ▼
 5. Requests Assessment
        │
        ▼
 6. Evidence Validated
        │
        ▼
 7. Financial Profile Built
        │
        ▼
 8. Trust Assessment Generated
        │
        ▼
 9. Fraud Analysis Completed
        │
        ▼
10. Explanation Generated
        │
        ▼
11. Recommendation Generated
        │
        ▼
12. Loan Officer Reviews
        │
        ▼
13. Decision Recorded
```

### Step-by-step breakdown

1. **Applicant** enters the system.
2. **Registers** — creates an account (triggers `Applicant Registered`).
3. **Grants Consent** — authorizes data usage (triggers `Consent Granted`).
4. **Uploads Financial Evidence** — submits supporting financial documents/data.
5. **Requests Assessment** — formally requests a trust assessment (triggers `Applicant Requested Assessment`).
6. **Evidence Validated** — submitted evidence is checked for validity/completeness.
7. **Financial Profile Built** — validated evidence is aggregated into a financial profile.
8. **Trust Assessment Generated** — a trust assessment is created from the financial profile (triggers `Trust Assessment Generated`).
9. **Fraud Analysis Completed** — fraud checks run against the applicant/profile (triggers `Fraud Analysis Completed`, and possibly `Fraud Alert Raised` if suspicious activity is found).
10. **Explanation Generated** — a human-readable rationale is produced for the assessment outcome.
11. **Recommendation Generated** — a recommended decision (approve/decline/review) is produced.
12. **Loan Officer Reviews** — a human reviewer examines the assessment, explanation, and recommendation.
13. **Decision Recorded** — the final decision is logged, closing out the workflow.

> Note: **Assessment Put On Hold** can occur at multiple points in this flow (e.g., between steps 6–9) if evidence is insufficient or additional review is required before proceeding.

---

## 3. Summary Flow (High-Level)

A condensed view of how the core entities relate to one another:

```
Applicant
    │
    ▼
Financial Evidence
    │
    ▼
Financial Profile
    │
    ▼
Trust Assessment
    │
    ├── Recommendation
    ├── Explanation
    └── Fraud Alert
```

**Reading this diagram:**
- An **Applicant** submits **Financial Evidence**.
- That evidence is aggregated into a **Financial Profile**.
- The profile feeds into a **Trust Assessment**.
- The Trust Assessment produces three associated outputs: a **Recommendation**, an **Explanation**, and — if warranted — a **Fraud Alert**.
