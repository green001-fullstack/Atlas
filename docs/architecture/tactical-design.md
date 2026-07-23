# Domain Model: Entity vs. Value Object Classification

This document classifies each core concept in the domain model as either an **Entity** or a **Value Object**, following standard Domain-Driven Design (DDD) principles, with rationale for each classification.

## Classification Principles

| Type | Defining Trait |
|---|---|
| **Entity** | Has a persistent, unique identity that endures over time, even if its attributes change. Two entities with identical attributes are still distinct if their identities differ. |
| **Value Object** | Defined entirely by its attributes/value. Immutable. Two value objects with identical attributes are considered equal and interchangeable — it has no identity of its own. |

---

## Classification Table

| Domain Concept | Entity or Value Object? | Why |
|---|---|---|
| Applicant | Entity | Has a persistent or unique identity that must be tracked across the system regardless of attribute changes (e.g., name, address updates). |
| Organization | Entity | Exists independently over time, tracked by identity rather than by its current attributes. |
| Organization User | Entity | Represents a specific person within an organization, with a distinct identity separate from others. |
| Assessment Request | Entity | Each request is unique and must be individually tracked, even if two requests have identical input data. |
| Trust Assessment | Entity | Has its own lifecycle and history (created, updated, finalized) that must be preserved over time. |
| Financial Evidence | Entity | Each piece of evidence is unique and traceable to its source, even if the content is duplicated elsewhere. |
| Fraud Alert | Entity | Can be opened, investigated, and resolved — it has a lifecycle and state transitions tied to its identity. |
| **Recommendation** | **Value Object** | Represents an immutable output/conclusion at a point in time. Two recommendations with the same content and rationale are interchangeable — it's defined by its value, not tracked by identity. |
| **Explanation** | **Value Object** | A descriptive artifact (e.g., reasoning text, factors considered) attached to a decision. It has no identity of its own — it's fully defined by its content. |
| **Financial Profile** | **Entity** | Represents an evolving aggregation of an applicant's financial standing over time. It persists, accumulates evidence, and is updated — making identity (tied to the applicant) meaningful. |
| **Trust Score** | **Value Object** | A calculated numeric/qualitative value at a point in time. Two trust scores with the same value are equal — recalculating produces a new value object, not a mutation of an existing one. |
| **Confidence Score** | **Value Object** | Same reasoning as Trust Score — a derived value with no independent identity; fully defined by its numeric value. |
| **Email Address** | **Value Object** | A classic DDD value object — defined entirely by its string value, immutable, and interchangeable if two addresses match exactly. |
| **Phone Number** | **Value Object** | Same reasoning as Email Address — defined by its value, immutable, and equal to any other phone number with the same digits/format. |

---

## Summary by Category

### Entities (have identity, persist over time)
- Applicant
- Organization
- Organization User
- Assessment Request
- Trust Assessment
- Financial Evidence
- Fraud Alert
- Financial Profile

### Value Objects (defined by value, immutable, no identity)
- Recommendation
- Explanation
- Trust Score
- Confidence Score
- Email Address
- Phone Number

---

## Notes & Edge Cases

- **Financial Profile** is the one concept with a plausible argument either way. If it's treated as a *snapshot* generated fresh for each assessment (never updated in place), it could instead be modeled as a **Value Object** owned by a `Trust Assessment` or `Financial Evidence` collection. It's classified here as an **Entity** under the assumption that it evolves and accumulates data over the applicant's lifetime rather than being recreated each time. Worth confirming against actual system behavior.
- **Recommendation** and **Explanation** are often modeled as value objects *nested inside* an Entity (e.g., as part of a `Trust Assessment`), rather than as standalone aggregates with their own repository/table identity.
- **Trust Score** and **Confidence Score** should typically be treated as **immutable** once created — any recalculation should produce a new value object rather than mutating the old one, to preserve historical accuracy in `Trust Assessment` history.
