package repository

import (
	"context"
	"booking/domain/entity"
)

// UserRepository defines the interface for user data operations
// This follows the Repository pattern and Dependency Inversion Principle
type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	GetByID(ctx context.Context, id uint) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	List(ctx context.Context, filter *entity.UserFilter) ([]*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id uint) error
	Count(ctx context.Context, filter *entity.UserFilter) (int64, error)
}

