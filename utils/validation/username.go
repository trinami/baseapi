package validation

import(
	"errors"
	"regexp"
)

func ValidateUsername(username string) (bool, error) {
	if len(username) < 3 || len(username) > 15 {
		return false, errors.New("username must be between 3 and 15 characters long")
	}

	var validUsername = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	if !validUsername.MatchString(username) {
		return false, errors.New("username can only contain letters, numbers, and underscores")
	}

	return true, nil
}
