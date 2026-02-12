package observer

import (
	"sync"
	"booking/domain/entity"
)

// EventType represents different types of events
type EventType string

const (
	UserCreated EventType = "user.created"
	UserUpdated EventType = "user.updated"
	UserDeleted EventType = "user.deleted"
)

// Event represents an event in the system
type Event struct {
	Type EventType
	Data interface{}
}

// Observer defines the interface for event observers
// Observer Pattern: Allows objects to be notified of changes
type Observer interface {
	Update(event Event)
}

// Subject manages observers and notifies them of events
type Subject struct {
	observers []Observer
	mu        sync.RWMutex
}

// NewSubject creates a new Subject
func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

// Attach adds an observer to the subject
func (s *Subject) Attach(observer Observer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.observers = append(s.observers, observer)
}

// Detach removes an observer from the subject
func (s *Subject) Detach(observer Observer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	for i, obs := range s.observers {
		if obs == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

// Notify notifies all observers of an event
func (s *Subject) Notify(event Event) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	for _, observer := range s.observers {
		go observer.Update(event) // Async notification
	}
}

// UserEventLogger is a concrete observer that logs user events
type UserEventLogger struct{}

// NewUserEventLogger creates a new UserEventLogger
func NewUserEventLogger() *UserEventLogger {
	return &UserEventLogger{}
}

// Update implements the Observer interface
func (l *UserEventLogger) Update(event Event) {
	switch event.Type {
	case UserCreated:
		if user, ok := event.Data.(*entity.User); ok {
			println("📝 [LOG] User created:", user.Username, "- Email:", user.Email)
		}
	case UserUpdated:
		if user, ok := event.Data.(*entity.User); ok {
			println("📝 [LOG] User updated:", user.Username)
		}
	case UserDeleted:
		if id, ok := event.Data.(uint); ok {
			println("📝 [LOG] User deleted: ID", id)
		}
	}
}

// UserEventNotifier is another concrete observer for notifications
type UserEventNotifier struct{}

// NewUserEventNotifier creates a new UserEventNotifier
func NewUserEventNotifier() *UserEventNotifier {
	return &UserEventNotifier{}
}

// Update implements the Observer interface
func (n *UserEventNotifier) Update(event Event) {
	switch event.Type {
	case UserCreated:
		if user, ok := event.Data.(*entity.User); ok {
			println("📧 [NOTIFY] Welcome email sent to:", user.Email)
		}
	case UserUpdated:
		if user, ok := event.Data.(*entity.User); ok {
			println("📧 [NOTIFY] Update notification sent to:", user.Email)
		}
	}
}

