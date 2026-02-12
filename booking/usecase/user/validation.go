package user

import (
	"errors"
	"regexp"
	"booking/domain/entity"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// validateUser validates user data
func (uc *userUseCase) validateUser(user *entity.User) error {
	if user.Email == "" {
		return errors.New("email is required")
	}
	
	if user.Username == "" {
		return errors.New("username is required")
	}
	
	if user.Password == "" {
		return errors.New("password is required")
	}
	
	// Email validation
	if uc.options.ValidateEmail && !emailRegex.MatchString(user.Email) {
		return errors.New("invalid email format")
	}
	
	// Password validation
	if uc.options.ValidatePassword {
		if len(user.Password) < uc.options.MinPasswordLen {
			return errors.New("password is too short")
		}
		if len(user.Password) > uc.options.MaxPasswordLen {
			return errors.New("password is too long")
		}
	}
	
	return nil
}

