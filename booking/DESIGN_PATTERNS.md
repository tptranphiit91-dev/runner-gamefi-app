# Design Patterns Implementation Guide

Tài liệu này giải thích chi tiết cách các Design Patterns được implement trong Booking Service.

## 📋 Table of Contents
1. [Singleton Pattern](#1-singleton-pattern)
2. [Factory Pattern](#2-factory-pattern)
3. [Strategy Pattern](#3-strategy-pattern)
4. [Observer Pattern](#4-observer-pattern)
5. [Functional Options Pattern](#5-functional-options-pattern)

---

## 1. Singleton Pattern

### 📍 Location
`infrastructure/database/postgres.go`

### 🎯 Purpose
Đảm bảo chỉ có **một instance duy nhất** của database connection trong toàn bộ application.

### 💡 Implementation

```go
var (
    instance *Database
    once     sync.Once
)

func GetInstance(config *Config) (*Database, error) {
    var err error
    
    once.Do(func() {
        // Initialize database connection only once
        db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if dbErr != nil {
            err = dbErr
            return
        }
        instance = &Database{DB: db}
    })
    
    return instance, err
}
```

### ✅ Benefits
- **Thread-safe**: Sử dụng `sync.Once` đảm bảo thread safety
- **Resource efficiency**: Tránh tạo nhiều connections không cần thiết
- **Global access**: Dễ dàng truy cập database từ mọi nơi

### 🔍 Usage Example
```go
db, err := database.GetInstance(dbConfig)
// Mọi lần gọi GetInstance đều trả về cùng một instance
```

---

## 2. Factory Pattern

### 📍 Locations
1. `delivery/http/handler/handler_factory.go` - Handler Factory
2. `infrastructure/database/database_factory.go` - Database Factory (NEW!)

### 🎯 Purpose
Tạo objects mà không cần expose logic khởi tạo phức tạp. Cho phép chọn implementation dựa trên configuration.

### 💡 Implementation

#### Handler Factory
```go
type HandlerFactory struct {
    userUseCase user.UserUseCase
}

func (f *HandlerFactory) CreateHandler(handlerType HandlerType) interface{} {
    switch handlerType {
    case UserHandlerType:
        return NewUserHandler(f.userUseCase)
    // Có thể thêm các handler types khác
    default:
        return nil
    }
}
```

#### Database Factory (NEW!)
```go
type DatabaseFactory struct {
    config  *config.Config
    subject *observer.Subject
}

func (f *DatabaseFactory) CreateUserRepository() (repository.UserRepository, error) {
    switch f.config.DatabaseType {
    case config.PostgresDB:
        return f.createPostgresUserRepository()
    case config.MongoDB:
        return f.createMongoUserRepository()
    default:
        return nil, fmt.Errorf("unsupported database type")
    }
}
```

### ✅ Benefits
- **Encapsulation**: Ẩn logic khởi tạo phức tạp
- **Flexibility**: Dễ dàng thêm handler/database types mới
- **Centralized creation**: Tất cả objects được tạo ở một nơi
- **Database abstraction**: Chuyển đổi giữa PostgreSQL và MongoDB dễ dàng

### 🔍 Usage Examples

**Handler Factory:**
```go
handlerFactory := handler.NewHandlerFactory(userUseCase)
userHandler := handlerFactory.GetUserHandler()
```

**Database Factory:**
```go
dbFactory := database.NewDatabaseFactory(cfg, subject)
userRepo, err := dbFactory.CreateUserRepository()
// Tự động chọn PostgreSQL hoặc MongoDB dựa trên config
```

---

## 3. Strategy Pattern

### 📍 Location
`usecase/user/password_strategy.go`

### 🎯 Purpose
Cho phép thay đổi thuật toán hash password tại runtime mà không thay đổi code sử dụng nó.

### 💡 Implementation

```go
// Strategy Interface
type PasswordHasher interface {
    Hash(password string) (string, error)
    Compare(hashedPassword, password string) error
}

// Concrete Strategy 1: Bcrypt
type BcryptHasher struct {
    cost int
}

func (h *BcryptHasher) Hash(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
    return string(bytes), err
}

// Concrete Strategy 2: SHA256
type SHA256Hasher struct {
    salt string
}

func (h *SHA256Hasher) Hash(password string) (string, error) {
    hash := sha256.New()
    hash.Write([]byte(password + h.salt))
    return hex.EncodeToString(hash.Sum(nil)), nil
}
```

### ✅ Benefits
- **Interchangeable algorithms**: Dễ dàng switch giữa các thuật toán
- **Open/Closed Principle**: Mở cho extension, đóng cho modification
- **Testability**: Dễ dàng mock strategies trong tests

### 🔍 Usage Example
```go
// Sử dụng Bcrypt
passwordHasher := user.NewBcryptHasher(10)

// Hoặc sử dụng SHA256
// passwordHasher := user.NewSHA256Hasher("my-salt")

userUseCase := user.NewUserUseCase(userRepo, passwordHasher)
```

---

## 4. Observer Pattern

### 📍 Location
`infrastructure/observer/event.go`

### 🎯 Purpose
Cho phép các objects (observers) được thông báo tự động khi có events xảy ra.

### 💡 Implementation

```go
// Observer Interface
type Observer interface {
    Update(event Event)
}

// Subject manages observers
type Subject struct {
    observers []Observer
    mu        sync.RWMutex
}

func (s *Subject) Attach(observer Observer) {
    s.observers = append(s.observers, observer)
}

func (s *Subject) Notify(event Event) {
    for _, observer := range s.observers {
        go observer.Update(event) // Async notification
    }
}

// Concrete Observer 1: Logger
type UserEventLogger struct{}

func (l *UserEventLogger) Update(event Event) {
    switch event.Type {
    case UserCreated:
        println("User created:", user.Username)
    }
}

// Concrete Observer 2: Notifier
type UserEventNotifier struct{}

func (n *UserEventNotifier) Update(event Event) {
    switch event.Type {
    case UserCreated:
        println("Welcome email sent to:", user.Email)
    }
}
```

### ✅ Benefits
- **Loose coupling**: Subject không cần biết chi tiết về observers
- **Dynamic subscription**: Có thể attach/detach observers tại runtime
- **Multiple observers**: Nhiều observers có thể lắng nghe cùng một event

### 🔍 Usage Example
```go
subject := observer.NewSubject()

// Attach observers
logger := observer.NewUserEventLogger()
notifier := observer.NewUserEventNotifier()
subject.Attach(logger)
subject.Attach(notifier)

// Notify all observers
subject.Notify(observer.Event{
    Type: observer.UserCreated,
    Data: user,
})
```

---

## 5. Functional Options Pattern

### 📍 Location
`usecase/user/user_usecase.go`

### 🎯 Purpose
Cung cấp cách linh hoạt để configure objects với optional parameters.

### 💡 Implementation

```go
// Options struct
type UseCaseOptions struct {
    ValidateEmail    bool
    ValidatePassword bool
    MinPasswordLen   int
    MaxPasswordLen   int
}

// Option function type
type UseCaseOption func(*UseCaseOptions)

// Option functions
func WithEmailValidation(validate bool) UseCaseOption {
    return func(o *UseCaseOptions) {
        o.ValidateEmail = validate
    }
}

func WithPasswordValidation(validate bool) UseCaseOption {
    return func(o *UseCaseOptions) {
        o.ValidatePassword = validate
    }
}

func WithPasswordLength(min, max int) UseCaseOption {
    return func(o *UseCaseOptions) {
        o.MinPasswordLen = min
        o.MaxPasswordLen = max
    }
}

// Constructor with variadic options
func NewUserUseCase(
    userRepo repository.UserRepository,
    passwordHasher PasswordHasher,
    opts ...UseCaseOption,
) UserUseCase {
    options := defaultOptions()
    
    for _, opt := range opts {
        opt(options)
    }
    
    return &userUseCase{
        userRepo:       userRepo,
        passwordHasher: passwordHasher,
        options:        options,
    }
}
```

### ✅ Benefits
- **Backward compatibility**: Thêm options mới không break existing code
- **Readable**: Self-documenting code
- **Flexible**: Có thể combine options theo nhiều cách
- **Default values**: Tự động có default values

### 🔍 Usage Example
```go
// Sử dụng với tất cả options
userUseCase := user.NewUserUseCase(
    userRepo,
    passwordHasher,
    user.WithEmailValidation(true),
    user.WithPasswordValidation(true),
    user.WithPasswordLength(8, 72),
)

// Hoặc chỉ một số options
userUseCase := user.NewUserUseCase(
    userRepo,
    passwordHasher,
    user.WithPasswordLength(10, 50),
)

// Hoặc không có options (sử dụng defaults)
userUseCase := user.NewUserUseCase(userRepo, passwordHasher)
```

---

## 🎓 Kết hợp các Patterns

Trong `cmd/api/main.go`, tất cả patterns được kết hợp:

```go
// 1. Singleton: Database connection
db, _ := database.GetInstance(dbConfig)

// 2. Observer: Event system
subject := observer.NewSubject()
subject.Attach(observer.NewUserEventLogger())
subject.Attach(observer.NewUserEventNotifier())

// 3. Strategy: Password hasher
passwordHasher := user.NewBcryptHasher(10)

// 4. Functional Options: UseCase configuration
userUseCase := user.NewUserUseCase(
    userRepo,
    passwordHasher,
    user.WithEmailValidation(true),
    user.WithPasswordLength(8, 72),
)

// 5. Factory: Handler creation
handlerFactory := handler.NewHandlerFactory(userUseCase)
```

---

## 📚 Tài liệu tham khảo

- [Gang of Four Design Patterns](https://en.wikipedia.org/wiki/Design_Patterns)
- [Go Design Patterns](https://github.com/trekhleb/go-patterns)
- [Functional Options in Go](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)

