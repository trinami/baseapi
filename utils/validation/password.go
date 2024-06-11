package validation

import(
	"errors"
	"regexp"
)

func ValidatePassword(password string) (bool, error) {
	if len(password) < 8 {
		return false, errors.New("password must be at least 8 characters long")
	}

	var (
		hasUpper   = regexp.MustCompile(`[A-Z]`).MatchString
		hasNumber  = regexp.MustCompile(`[0-9]`).MatchString
		hasSpecial = regexp.MustCompile(`[!@#~$%^&*(),.?":{}|<>]`).MatchString
	)

	if !hasUpper(password) {
		return false, errors.New("password must contain at least one uppercase letter")
	}

	if !hasNumber(password) {
		return false, errors.New("password must contain at least one number")
	}

	if !hasSpecial(password) {
		return false, errors.New("password must contain at least one special character")
	}

	return true, nil
}
