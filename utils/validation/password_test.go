package validation

import (
	"testing"
)

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		password string
		valid    bool
	}{
		{"Password1!", true},
		{"password1", false},
		{"Password", false},
		{"Password!", false},
		{"Pass1!", false},
		{"12345678", false},
		{"Password123!", true},
		{"Passw0rd!", true},
		{"Pa$$w0rd", true},
		{"", false},
	}

	for _, test := range tests {
		valid, _ := ValidatePassword(test.password)
		if valid != test.valid {
			t.Errorf("ValidatePassword(%q) = %v; want %v", test.password, valid, test.valid)
		}
	}
}
