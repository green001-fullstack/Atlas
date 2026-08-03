package domain

import (
	"time"
)

type Consent struct {
	grantedAt   time.Time
    withdrawnAt *time.Time
    version     string
}

func NewConsent( version string, grantedAt time.Time) (*Consent, error){
	// Constructors should never create invalid domain objects
	if version == "" {
		return nil, ErrInvalidConsent
	}

	if grantedAt.IsZero() {
		return nil, ErrInvalidConsent
	}

	return &Consent{
		grantedAt: grantedAt,
		version: version,
	}, nil
}

func (c *Consent) Version() (string){
	 return c.version
}

func (c *Consent) GrantedAt() time.Time {
	return c.grantedAt
}

func (c *Consent) WithdrawnAt() (time.Time, bool) {
	if c.withdrawnAt == nil {
		return time.Time{}, false
	}

	return *c.withdrawnAt, true
}

func (c *Consent) IsActive()bool {
	return c.withdrawnAt == nil
}