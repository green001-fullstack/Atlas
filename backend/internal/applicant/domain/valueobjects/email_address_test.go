
package valueobjects
 
import "testing"
 
func TestNewEmailAddress(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantValue string // expected normalized value, only checked when wantErr is false
		wantErr   bool
	}{
		{
			name:      "valid email",
			input:     "user@example.com",
			wantValue: "user@example.com",
			wantErr:   false,
		},
		{
			name:    "empty email",
			input:   "",
			wantErr: true,
		},
		{
			name:    "missing @ symbol",
			input:   "userexample.com",
			wantErr: true,
		},
		{
			name:    "missing domain",
			input:   "user@",
			wantErr: true,
		},
		{
			name:    "missing local part",
			input:   "@example.com",
			wantErr: true,
		},
		{
			name:      "leading and trailing spaces are trimmed",
			input:     "  user@example.com  ",
			wantValue: "user@example.com",
			wantErr:   false,
		},
		{
			name:      "uppercase is normalized to lowercase",
			input:     "User@Example.COM",
			wantValue: "user@example.com",
			wantErr:   false,
		},
		{
			name:    "multiple @ symbols",
			input:   "user@@example.com",
			wantErr: true,
		},
	}
 
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEmailAddress(tt.input)
 
			if tt.wantErr {
				if err == nil {
					t.Fatalf("NewEmailAddress(%q) expected an error, got nil", tt.input)
				}
				if err != ErrInvalidEmailAddress {
					t.Errorf("NewEmailAddress(%q) error = %v, want %v", tt.input, err, ErrInvalidEmailAddress)
				}
				return
			}
 
			if err != nil {
				t.Fatalf("NewEmailAddress(%q) unexpected error: %v", tt.input, err)
			}
			if got.String() != tt.wantValue {
				t.Errorf("NewEmailAddress(%q).String() = %q, want %q", tt.input, got.String(), tt.wantValue)
			}
		})
	}
}