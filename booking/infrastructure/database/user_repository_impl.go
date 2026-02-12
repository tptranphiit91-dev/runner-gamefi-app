package database

import (
	"context"
	"booking/domain/entity"
	"booking/domain/repository"
	"booking/infrastructure/observer"
	
	"gorm.io/gorm"
)

// userRepositoryImpl implements the UserRepository interface
type userRepositoryImpl struct {
	db      *gorm.DB
	subject *observer.Subject
}

// NewUserRepository creates a new user repository
// This is a Factory function
func NewUserRepository(db *gorm.DB, subject *observer.Subject) repository.UserRepository {
	return &userRepositoryImpl{
		db:      db,
		subject: subject,
	}
}

// Create creates a new user
func (r *userRepositoryImpl) Create(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	
	// Notify observers
	r.subject.Notify(observer.Event{
		Type: observer.UserCreated,
		Data: user,
	})
	
	return nil
}

// GetByID retrieves a user by ID
func (r *userRepositoryImpl) GetByID(ctx context.Context, id uint) (*entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail retrieves a user by email
func (r *userRepositoryImpl) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByUsername retrieves a user by username
func (r *userRepositoryImpl) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// List retrieves users based on filter
func (r *userRepositoryImpl) List(ctx context.Context, filter *entity.UserFilter) ([]*entity.User, error) {
	var users []*entity.User
	query := r.db.WithContext(ctx)
	
	if filter != nil {
		if filter.Email != nil {
			query = query.Where("email = ?", *filter.Email)
		}
		if filter.Username != nil {
			query = query.Where("username = ?", *filter.Username)
		}
		if filter.IsActive != nil {
			query = query.Where("is_active = ?", *filter.IsActive)
		}
		if filter.Limit > 0 {
			query = query.Limit(filter.Limit)
		}
		if filter.Offset > 0 {
			query = query.Offset(filter.Offset)
		}
	}
	
	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	
	return users, nil
}

// Update updates a user
func (r *userRepositoryImpl) Update(ctx context.Context, user *entity.User) error {
	if err := r.db.WithContext(ctx).Save(user).Error; err != nil {
		return err
	}
	
	// Notify observers
	r.subject.Notify(observer.Event{
		Type: observer.UserUpdated,
		Data: user,
	})
	
	return nil
}

// Delete deletes a user by ID
func (r *userRepositoryImpl) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&entity.User{}, id).Error; err != nil {
		return err
	}
	
	// Notify observers
	r.subject.Notify(observer.Event{
		Type: observer.UserDeleted,
		Data: id,
	})
	
	return nil
}

// Count counts users based on filter
func (r *userRepositoryImpl) Count(ctx context.Context, filter *entity.UserFilter) (int64, error) {
	var count int64
	query := r.db.WithContext(ctx).Model(&entity.User{})
	
	if filter != nil {
		if filter.Email != nil {
			query = query.Where("email = ?", *filter.Email)
		}
		if filter.Username != nil {
			query = query.Where("username = ?", *filter.Username)
		}
		if filter.IsActive != nil {
			query = query.Where("is_active = ?", *filter.IsActive)
		}
	}
	
	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}
	
	return count, nil
}

