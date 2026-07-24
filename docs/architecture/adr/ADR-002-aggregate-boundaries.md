Decision 1:

Atlas will enforce business consistency through small, business-focused aggregates rather than large object graphs.

Reason:

Small aggregates improve consistency, scalability, and maintainability while aligning with Domain-Driven Design principles.

Decision 2:

Assessment Requests are modeled as a separate Aggregate from Trust Assessments because requesting an assessment and producing an assessment are distinct business processes with independent life cycles.