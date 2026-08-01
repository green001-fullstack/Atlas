package valueobjects

import (
	"errors"
	"testing"
)

func TestNormalizeName(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "trims leading and trailing spaces",
			input: "   John Doe   ",
			want:  "John Doe",
		},
		{
			name:  "collapses multiple internal spaces",
			input: "John     Doe",
			want:  "John Doe",
		},
		{
			name:  "collapses tabs and newlines into single spaces",
			input: "John\t\nDoe",
			want:  "John Doe",
		},
		{
			name:  "returns empty string for empty input",
			input: "",
			want:  "",
		},
		{
			name:  "returns empty string for whitespace-only input",
			input: "     ",
			want:  "",
		},
		{
			name:  "leaves already-clean string unchanged",
			input: "Jane Smith",
			want:  "Jane Smith",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizeName(tt.input)
			if got != tt.want {
				t.Errorf("normalizeName(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestValidateName(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr error
	}{
		{
			name:    "valid simple name",
			input:   "John Doe",
			wantErr: nil,
		},
		{
			name:    "empty string is invalid",
			input:   "",
			wantErr: ErrInvalidName,
		},
		{
			name:    "name containing digits is invalid",
			input:   "John123",
			wantErr: ErrInvalidName,
		},
		{
			name:    "name with a single digit is invalid",
			input:   "Agent007",
			wantErr: ErrInvalidName,
		},
		{
			name:    "name exactly at 100 characters is valid",
			input:   strRepeat("a", 100),
			wantErr: nil,
		},
		{
			name:    "name exceeding 100 characters is invalid",
			input:   strRepeat("a", 101),
			wantErr: ErrInvalidName,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateName(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("validateName(%q) error = %v, want %v", tt.input, err, tt.wantErr)
			}
		})
	}
}

func TestNewName(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantValue string
		wantErr   bool
	}{
		{
			name:      "valid name is created and normalized",
			input:     "  John   Doe  ",
			wantValue: "John Doe",
			wantErr:   false,
		},
		{
			name:      "empty input returns error",
			input:     "",
			wantErr:   true,
		},
		{
			name:      "whitespace-only input returns error",
			input:     "     ",
			wantErr:   true,
		},
		{
			name:      "name with digits returns error",
			input:     "John5",
			wantErr:   true,
		},
		{
			name:      "name exceeding max length returns error",
			input:     strRepeat("a", 150),
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewName(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Fatalf("NewName(%q) expected error, got nil", tt.input)
				}
				return
			}

			if err != nil {
				t.Fatalf("NewName(%q) unexpected error: %v", tt.input, err)
			}

			if got.String() != tt.wantValue {
				t.Errorf("NewName(%q).String() = %q, want %q", tt.input, got.String(), tt.wantValue)
			}
		})
	}
}

// strRepeat is a small test helper to build strings of a given length
// without needing to import "strings" purely for repetition in test data.
func strRepeat(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}