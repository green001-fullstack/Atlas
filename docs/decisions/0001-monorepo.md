# ADR-0001: Adopt a Monorepo

## Status

Accepted

## Context

Atlas consists of multiple components including the backend, frontend, documentation, deployment configuration, and automation scripts.

## Decision

The project will use a monorepo to keep all components together.

## Consequences

### Positive

- Easier onboarding
- Single source of truth
- Simpler CI/CD
- Easier coordinated changes

### Negative

- Larger repository
- CI pipelines may become more complex as the project grows