# 💡 Code Examples & Use Cases

Các ví dụ cụ thể về cách sử dụng và mở rộng code.

---

## 1. Thay đổi Password Hashing Strategy

### Hiện tại (Bcrypt)
```go
// cmd/api/main.go
passwordHasher := user.NewBcryptHasher(10)
```

### Thay đổi sang SHA256
```go
// cmd/api/main.go
passwordHasher := user.NewSHA256Hasher("my-secret-salt")
```

### Tạo Strategy mới (MD5 - for demo only)
```go
// usecase/user/password_strategy.go
type MD5Hasher struct {
    salt string
}

func NewMD5Hasher(salt string) *MD5Hasher {
    return &MD5Hasher{salt: salt}
}

func (h *MD5Hasher) Hash(password string) (string, error) {
    hash := md5.New()
    hash.Write([]byte(password + h.salt))
    return hex.EncodeToString(hash.Sum(nil)), nil
}

func (h *MD5Hasher) Compare(hashedPassword, password string) error {
    newHash, err := h.Hash(password)
    if err != nil {
        return err
    }
    if newHash != hashedPassword {
        return errors.New("password mismatch")
    }
    return nil
}
```

---

## 2. Thêm Observer mới

### Email Observer
```go
// infrastructure/observer/event.go
type UserEmailObserver struct {
    emailService EmailService
}

func NewUserEmailObserver(emailService EmailService) *UserEmailObserver {
    return &UserEmailObserver{emailService: emailService}
}

func (o *UserEmailObserver) Update(event Event) {
    switch event.Type {
    case UserCreated:
        if user, ok := event.Data.(*entity.User); ok {
            o.emailService.SendWelcomeEmail(user.Email, user.FullName)
        }
    }
}
```

### Sử dụng
```go
// cmd/api/main.go
emailObserver := observer.NewUserEmailObserver(emailService)
subject.Attach(emailObserver)
```

---

## 3. Thêm Functional Options mới

### Thêm option cho username validation
```go
// usecase/user/user_usecase.go
type UseCaseOptions struct {
    ValidateEmail    bool
    ValidatePassword bool
    ValidateUsername bool  // NEW
    MinPasswordLen   int
    MaxPasswordLen   int
    MinUsernameLen   int   // NEW
}

func WithUsernameValidation(validate bool, minLen int) UseCaseOption {
    return func(o *UseCaseOptions) {
        o.ValidateUsername = validate
        o.MinUsernameLen = minLen
    }
}
```

### Sử dụng
```go
userUseCase := user.NewUserUseCase(
    userRepo,
    passwordHasher,
    user.WithEmailValidation(true),
    user.WithPasswordLength(8, 72),
    user.WithUsernameValidation(true, 3), // NEW
)
```

---

## 4. Thêm Handler mới với Factory

### Tạo Booking Handler
```go
// delivery/http/handler/booking_handler.go
type BookingHandler struct {
    bookingUseCase booking.BookingUseCase
}

func NewBookingHandler(bookingUseCase booking.BookingUseCase) *BookingHandler {
    return &BookingHandler{bookingUseCase: bookingUseCase}
}
```

### Cập nhật Factory
```go
// delivery/http/handler/handler_factory.go
const (
    UserHandlerType    HandlerType = "user"
    BookingHandlerType HandlerType = "booking" // NEW
)

type HandlerFactory struct {
    userUseCase    user.UserUseCase
    bookingUseCase booking.BookingUseCase // NEW
}

func (f *HandlerFactory) CreateHandler(handlerType HandlerType) interface{} {
    switch handlerType {
    case UserHandlerType:
        return NewUserHandler(f.userUseCase)
    case BookingHandlerType:
        return NewBookingHandler(f.bookingUseCase) // NEW
    default:
        return nil
    }
}
```

---

## 5. API Request Examples

### Create User
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "username": "johndoe",
    "password": "SecurePass123!",
    "full_name": "John Doe",
    "phone": "+1234567890"
  }'
```

### List Users with Filters
```bash
# Get active users only
curl "http://localhost:8080/api/v1/users?is_active=true"

# Get with pagination
curl "http://localhost:8080/api/v1/users?limit=10&offset=0"

# Search by email
curl "http://localhost:8080/api/v1/users?email=john@example.com"
```

### Update User
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "full_name": "John Updated",
    "phone": "+9999999999",
    "is_active": false
  }'
```

---

## 6. Testing Examples

### Unit Test cho Password Strategy
```go
// usecase/user/password_strategy_test.go
func TestBcryptHasher(t *testing.T) {
    hasher := NewBcryptHasher(10)
    
    password := "testpassword123"
    hashed, err := hasher.Hash(password)
    
    assert.NoError(t, err)
    assert.NotEmpty(t, hashed)
    
    // Test compare
    err = hasher.Compare(hashed, password)
    assert.NoError(t, err)
    
    // Test wrong password
    err = hasher.Compare(hashed, "wrongpassword")
    assert.Error(t, err)
}
```

### Integration Test
```go
// delivery/http/handler/user_handler_test.go
func TestCreateUser(t *testing.T) {
    // Setup
    router := setupTestRouter()
    
    // Request
    body := `{
        "email": "test@example.com",
        "username": "testuser",
        "password": "password123"
    }`
    
    req, _ := http.NewRequest("POST", "/api/v1/users", strings.NewReader(body))
    req.Header.Set("Content-Type", "application/json")
    
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    
    // Assert
    assert.Equal(t, http.StatusCreated, w.Code)
}
```

---

## 7. Extending the System

### Thêm Booking Entity
```go
// domain/entity/booking.go
type Booking struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    UserID    uint      `json:"user_id" gorm:"not null"`
    User      User      `json:"user" gorm:"foreignKey:UserID"`
    ServiceID uint      `json:"service_id"`
    StartTime time.Time `json:"start_time"`
    EndTime   time.Time `json:"end_time"`
    Status    string    `json:"status"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

### Thêm Booking Repository
```go
// domain/repository/booking_repository.go
type BookingRepository interface {
    Create(ctx context.Context, booking *entity.Booking) error
    GetByID(ctx context.Context, id uint) (*entity.Booking, error)
    GetByUserID(ctx context.Context, userID uint) ([]*entity.Booking, error)
    Update(ctx context.Context, booking *entity.Booking) error
    Delete(ctx context.Context, id uint) error
}
```

---

## 8. Environment Configuration

### Development
```env
SERVER_PORT=8080
DB_HOST=localhost
DB_NAME=booking_dev
```

### Production
```env
SERVER_PORT=80
DB_HOST=prod-db.example.com
DB_NAME=booking_prod
DB_SSLMODE=require
```

### Testing
```env
SERVER_PORT=8081
DB_HOST=localhost
DB_NAME=booking_test
```

---

## 🎯 Best Practices Demonstrated

1. **Separation of Concerns**: Mỗi layer có trách nhiệm riêng
2. **Dependency Injection**: Dependencies được inject qua constructor
3. **Interface-based Design**: Program to interfaces, not implementations
4. **Error Handling**: Proper error propagation
5. **Configuration**: Environment-based configuration
6. **Logging**: Structured logging với Observer pattern
7. **Validation**: Input validation ở use case layer
8. **Security**: Password hashing, no password in responses

---

Happy Coding! 🚀

