package domain

import "errors"

var ErrInvalidConsent = errors.New("invalid consent")
var ErrConsentNotGranted = errors.New("consent not granted")
var ErrConsentAlreadyWithdrawn = errors.New("consent already withdrawn")
var ErrAssessmentAlreadyRequested = errors.New("assessment already withdrawn")
var ErrInvalidApplicant = errors.New("invalid applicant")
var ErrConsentAlreadyGranted = errors.New("Consent already granted")
var ErrApplicantAlreadyExists = errors.New("applicant already exists")



