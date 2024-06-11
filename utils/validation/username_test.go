package validation

import (
	"testing"
)

func TestValidateUsername(t *testing.T) {
	tests := []struct {
		username string
		valid    bool
	}{
		{"valid_user123", true},
		{"invalid user", false},
		{"toolongusername12345", false},
		{"sh", false},
		{"validUser", true},
		{"invalid@user", false},
		{"validUser_123", true},
		{"", false},
		{"user", true},
		{"user-name", false},
	}

	for _, test := range tests {
		valid, _ := ValidateUsername(test.username)
		if valid != test.valid {
			t.Errorf("ValidateUsername(%q) = %v; want %v", test.username, valid, test.valid)
		}
	}
}
