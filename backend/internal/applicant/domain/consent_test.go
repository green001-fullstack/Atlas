package domain

import (
	"testing"
	"time"
)

func TestNewConsent_WithValidData(t *testing.T) {
	grantedAt := time.Date(2026, time.August, 3, 10, 0, 0, 0, time.UTC)

	consent, err := NewConsent("v1", grantedAt)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if consent == nil {
		t.Fatal("expected consent to be created")
	}

	if consent.Version() != "v1" {
		t.Errorf("expected version to be v1, got %s", consent.Version())
	}

	if !consent.GrantedAt().Equal(grantedAt) {
		t.Errorf("expected grantedAt to be %v, got %v", grantedAt, consent.GrantedAt())
	}

	if !consent.IsActive() {
		t.Error("expected consent to be active")
	}

	_, ok := consent.WithdrawnAt()
	if ok {
		t.Error("expected withdrawnAt to be empty")
	}
}

func TestNewConsent_WithEmptyVersion_ReturnsError(t *testing.T) {
	grantedAt := time.Now()

	consent, err := NewConsent("", grantedAt)

	if err != ErrInvalidConsent {
		t.Errorf("expected ErrInvalidConsent, got %v", err)
	}

	if consent != nil {
		t.Error("expected consent to be nil")
	}
}

func TestNewConsent_WithZeroGrantedAt_ReturnsError(t *testing.T) {
	consent, err := NewConsent("v1", time.Time{})

	if err != ErrInvalidConsent {
		t.Errorf("expected ErrInvalidConsent, got %v", err)
	}

	if consent != nil {
		t.Error("expected consent to be nil")
	}
}