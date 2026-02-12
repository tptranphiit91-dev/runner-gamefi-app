package handler

import (
	"booking/usecase/user"
)

// HandlerType represents different types of handlers
type HandlerType string

const (
	UserHandlerType HandlerType = "user"
)

// HandlerFactory creates handlers based on type
// Factory Pattern: Creates different types of handlers
type HandlerFactory struct {
	userUseCase user.UserUseCase
}

// NewHandlerFactory creates a new handler factory
func NewHandlerFactory(userUseCase user.UserUseCase) *HandlerFactory {
	return &HandlerFactory{
		userUseCase: userUseCase,
	}
}

// CreateHandler creates a handler based on type
func (f *HandlerFactory) CreateHandler(handlerType HandlerType) interface{} {
	switch handlerType {
	case UserHandlerType:
		return NewUserHandler(f.userUseCase)
	default:
		return nil
	}
}

// GetUserHandler returns a user handler
func (f *HandlerFactory) GetUserHandler() *UserHandler {
	return f.CreateHandler(UserHandlerType).(*UserHandler)
}

