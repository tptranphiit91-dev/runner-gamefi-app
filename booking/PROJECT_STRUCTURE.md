# 📁 Project Structure

## Tổng quan cấu trúc thư mục

```
booking/
├── cmd/                          # Application entry points
│   └── api/
│       └── main.go              # Main application (wires everything together)
│
├── config/                       # Configuration management
│   └── config.go                # Load config from environment variables
│
├── delivery/                     # Delivery Layer (Presentation)
│   └── http/
│       ├── handler/
│       │   ├── user_handler.go          # HTTP handlers for user endpoints
│       │   └── handler_factory.go       # Factory Pattern for handlers
│       ├── middleware/
│       │   ├── cors.go                  # CORS middleware
│       │   └── logger.go                # Logging middleware
│       └── router.go                    # Route configuration
│
├── usecase/                      # Use Case Layer (Business Logic)
│   └── user/
│       ├── user_usecase.go              # User business logic
│       ├── password_strategy.go         # Strategy Pattern for password hashing
│       └── validation.go                # Input validation
│
├── domain/                       # Domain Layer (Core Business)
│   ├── entity/
│   │   └── user.go                      # User entity
│   └── repository/
│       └── user_repository.go           # Repository interface
│
├── infrastructure/               # Infrastructure Layer (External)
│   ├── database/
│   │   ├── postgres.go                  # Singleton database connection
│   │   └── user_repository_impl.go      # Repository implementation
│   └── observer/
│       └── event.go                     # Observer Pattern for events
│
├── .env.example                  # Environment variables template
├── .gitignore                    # Git ignore rules
├── docker-compose.yml            # Docker Compose for PostgreSQL
├── Makefile                      # Build and run commands
├── go.mod                        # Go module dependencies
├── go.sum                        # Go module checksums
├── README.md                     # Main documentation
├── QUICKSTART.md                 # Quick start guide
├── DESIGN_PATTERNS.md            # Design patterns documentation
└── test_api.sh                   # API test script
```

---

## 📊 Layer Dependencies

```
Delivery Layer (HTTP)
    ↓ depends on
Use Case Layer (Business Logic)
    ↓ depends on
Domain Layer (Entities & Interfaces)
    ↑ implemented by
Infrastructure Layer (Database, External Services)
```

**Dependency Rule**: Các layer bên trong không biết gì về các layer bên ngoài.

---

## 🎯 Design Patterns Map

| Pattern | Location | Purpose |
|---------|----------|---------|
| **Singleton** | `infrastructure/database/postgres.go` | Single database instance |
| **Factory** | `delivery/http/handler/handler_factory.go` | Create handlers |
| **Strategy** | `usecase/user/password_strategy.go` | Interchangeable password hashers |
| **Observer** | `infrastructure/observer/event.go` | Event notifications |
| **Functional Options** | `usecase/user/user_usecase.go` | Flexible configuration |

---

## 📝 File Descriptions

### Domain Layer
- **`domain/entity/user.go`**
  - User entity với GORM tags
  - UserFilter cho query filtering
  
- **`domain/repository/user_repository.go`**
  - Interface định nghĩa contract cho user data access
  - CRUD operations

### Use Case Layer
- **`usecase/user/user_usecase.go`**
  - Business logic cho user operations
  - Sử dụng Functional Options Pattern
  - Validation và password hashing
  
- **`usecase/user/password_strategy.go`**
  - Strategy Pattern cho password hashing
  - BcryptHasher và SHA256Hasher implementations
  
- **`usecase/user/validation.go`**
  - Input validation logic
  - Email và password validation

### Infrastructure Layer
- **`infrastructure/database/postgres.go`**
  - Singleton Pattern cho database connection
  - Thread-safe với sync.Once
  - Auto migration
  
- **`infrastructure/database/user_repository_impl.go`**
  - Implementation của UserRepository interface
  - GORM operations
  - Observer notifications
  
- **`infrastructure/observer/event.go`**
  - Observer Pattern implementation
  - Subject, Observer interface
  - UserEventLogger và UserEventNotifier

### Delivery Layer
- **`delivery/http/handler/user_handler.go`**
  - HTTP request handlers
  - Request/Response DTOs
  - Error handling
  
- **`delivery/http/handler/handler_factory.go`**
  - Factory Pattern cho handlers
  - Centralized handler creation
  
- **`delivery/http/middleware/`**
  - CORS middleware
  - Logger middleware
  
- **`delivery/http/router.go`**
  - Route configuration
  - Middleware setup
  - Gin engine setup

### Configuration & Main
- **`config/config.go`**
  - Load configuration từ environment
  - Server và Database config
  
- **`cmd/api/main.go`**
  - Application entry point
  - Wire all components together
  - Demonstrate all design patterns

---

## 🔄 Request Flow

```
1. HTTP Request
   ↓
2. Router (delivery/http/router.go)
   ↓
3. Middleware (CORS, Logger)
   ↓
4. Handler (delivery/http/handler/user_handler.go)
   ↓
5. UseCase (usecase/user/user_usecase.go)
   ↓
6. Repository Interface (domain/repository/user_repository.go)
   ↓
7. Repository Implementation (infrastructure/database/user_repository_impl.go)
   ↓
8. Database (infrastructure/database/postgres.go)
   ↓
9. Observer Notification (infrastructure/observer/event.go)
   ↓
10. HTTP Response
```

---

## 🧩 Component Relationships

### Main Application Wire-up
```go
// 1. Database (Singleton)
db := database.GetInstance(config)

// 2. Observer (Observer Pattern)
subject := observer.NewSubject()
subject.Attach(logger)
subject.Attach(notifier)

// 3. Repository
userRepo := database.NewUserRepository(db, subject)

// 4. Password Hasher (Strategy)
hasher := user.NewBcryptHasher(10)

// 5. UseCase (Functional Options)
userUseCase := user.NewUserUseCase(
    userRepo, 
    hasher,
    user.WithEmailValidation(true),
)

// 6. Handler (Factory)
factory := handler.NewHandlerFactory(userUseCase)
userHandler := factory.GetUserHandler()

// 7. Router
router := http.NewRouter(factory)
```

---

## 📚 Key Files to Study

Nếu bạn muốn học Clean Architecture và Design Patterns, đọc theo thứ tự:

1. **`domain/entity/user.go`** - Hiểu entities
2. **`domain/repository/user_repository.go`** - Hiểu interfaces
3. **`infrastructure/database/postgres.go`** - Singleton Pattern
4. **`infrastructure/observer/event.go`** - Observer Pattern
5. **`usecase/user/password_strategy.go`** - Strategy Pattern
6. **`usecase/user/user_usecase.go`** - Functional Options Pattern
7. **`delivery/http/handler/handler_factory.go`** - Factory Pattern
8. **`cmd/api/main.go`** - Xem tất cả được kết hợp

---

## 🎓 Learning Path

1. **Beginner**: Đọc README.md và QUICKSTART.md
2. **Intermediate**: Đọc DESIGN_PATTERNS.md
3. **Advanced**: Đọc source code theo thứ tự trên
4. **Expert**: Thử thêm features mới (booking entity, authentication, etc.)

