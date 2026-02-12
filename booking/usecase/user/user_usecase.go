package user

import (
	"context"
	"errors"
	"booking/domain/entity"
	"booking/domain/repository"
	
	"gorm.io/gorm"
)

// UserUseCase defines the interface for user business logic
type UserUseCase interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUserByID(ctx context.Context, id uint) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	ListUsers(ctx context.Context, filter *entity.UserFilter) ([]*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, id uint) error
	CountUsers(ctx context.Context, filter *entity.UserFilter) (int64, error)
}

// userUseCase implements UserUseCase
type userUseCase struct {
	userRepo       repository.UserRepository
	passwordHasher PasswordHasher
	options        *UseCaseOptions
}

// UseCaseOptions holds optional configuration for the use case
// Functional Options Pattern: Allows flexible configuration
type UseCaseOptions struct {
	ValidateEmail    bool
	ValidatePassword bool
	MinPasswordLen   int
	MaxPasswordLen   int
}

// UseCaseOption is a function that configures UseCaseOptions
type UseCaseOption func(*UseCaseOptions)

// WithEmailValidation enables email validation
func WithEmailValidation(validate bool) UseCaseOption {
	return func(o *UseCaseOptions) {
		o.ValidateEmail = validate
	}
}

// WithPasswordValidation enables password validation
func WithPasswordValidation(validate bool) UseCaseOption {
	return func(o *UseCaseOptions) {
		o.ValidatePassword = validate
	}
}

// WithPasswordLength sets password length constraints
func WithPasswordLength(min, max int) UseCaseOption {
	return func(o *UseCaseOptions) {
		o.MinPasswordLen = min
		o.MaxPasswordLen = max
	}
}

// defaultOptions returns default use case options
func defaultOptions() *UseCaseOptions {
	return &UseCaseOptions{
		ValidateEmail:    true,
		ValidatePassword: true,
		MinPasswordLen:   8,
		MaxPasswordLen:   72,
	}
}

// NewUserUseCase creates a new user use case with functional options
// Functional Options Pattern implementation
func NewUserUseCase(
	userRepo repository.UserRepository,
	passwordHasher PasswordHasher,
	opts ...UseCaseOption,
) UserUseCase {
	options := defaultOptions()
	
	// Apply all options
	for _, opt := range opts {
		opt(options)
	}
	
	return &userUseCase{
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
		options:        options,
	}
}

// CreateUser creates a new user
func (uc *userUseCase) CreateUser(ctx context.Context, user *entity.User) error {
	// Validate input
	if err := uc.validateUser(user); err != nil {
		return err
	}
	
	// Check if user already exists
	existingUser, err := uc.userRepo.GetByEmail(ctx, user.Email)
	if err == nil && existingUser != nil {
		return errors.New("user with this email already exists")
	}
	
	existingUser, err = uc.userRepo.GetByUsername(ctx, user.Username)
	if err == nil && existingUser != nil {
		return errors.New("user with this username already exists")
	}
	
	// Hash password
	hashedPassword, err := uc.passwordHasher.Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	
	// Create user
	return uc.userRepo.Create(ctx, user)
}

// GetUserByID retrieves a user by ID
func (uc *userUseCase) GetUserByID(ctx context.Context, id uint) (*entity.User, error) {
	return uc.userRepo.GetByID(ctx, id)
}

// GetUserByEmail retrieves a user by email
func (uc *userUseCase) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	return uc.userRepo.GetByEmail(ctx, email)
}

// GetUserByUsername retrieves a user by username
func (uc *userUseCase) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	return uc.userRepo.GetByUsername(ctx, username)
}

// ListUsers retrieves users based on filter
func (uc *userUseCase) ListUsers(ctx context.Context, filter *entity.UserFilter) ([]*entity.User, error) {
	return uc.userRepo.List(ctx, filter)
}

// UpdateUser updates a user
func (uc *userUseCase) UpdateUser(ctx context.Context, user *entity.User) error {
	// Check if user exists
	existingUser, err := uc.userRepo.GetByID(ctx, user.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}
	
	// If password is being updated, hash it
	if user.Password != "" && user.Password != existingUser.Password {
		hashedPassword, err := uc.passwordHasher.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = hashedPassword
	} else {
		user.Password = existingUser.Password
	}
	
	return uc.userRepo.Update(ctx, user)
}

// DeleteUser deletes a user
func (uc *userUseCase) DeleteUser(ctx context.Context, id uint) error {
	return uc.userRepo.Delete(ctx, id)
}

// CountUsers counts users based on filter
func (uc *userUseCase) CountUsers(ctx context.Context, filter *entity.UserFilter) (int64, error) {
	return uc.userRepo.Count(ctx, filter)
}

