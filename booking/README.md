# Booking Service - Clean Architecture with Design Patterns

Service backend được xây dựng với Gin framework (Golang), tuân theo Clean Architecture và áp dụng các Design Patterns.

**✨ Hỗ trợ cả PostgreSQL và MongoDB databases!**

> 📚 **New here?** Start with [INDEX.md](INDEX.md) for a guided tour of all documentation!

## 🏗️ Kiến trúc

### Clean Architecture Layers

```
booking/
├── cmd/api/              # Application entry point
├── config/               # Configuration management
├── delivery/             # Delivery Layer (HTTP handlers, middleware)
│   └── http/
│       ├── handler/      # HTTP request handlers
│       ├── middleware/   # HTTP middleware
│       └── router.go     # Route configuration
├── usecase/              # Use Case Layer (Business logic)
│   └── user/
├── domain/               # Domain Layer (Entities, Interfaces)
│   ├── entity/           # Domain entities
│   └── repository/       # Repository interfaces
└── infrastructure/       # Infrastructure Layer (Database, External services)
    ├── database/         # Database implementation
    └── observer/         # Observer pattern implementation
```

## 🎯 Design Patterns

### 1. **Singleton Pattern**
- **File**: `infrastructure/database/postgres.go`
- **Mục đích**: Đảm bảo chỉ có một instance của database connection
- **Implementation**: Sử dụng `sync.Once` để thread-safe initialization

### 2. **Factory Pattern**
- **Files**:
  - `delivery/http/handler/handler_factory.go` - Handler creation
  - `infrastructure/database/database_factory.go` - Database selection (NEW!)
- **Mục đích**: Tạo objects dựa trên type/configuration
- **Implementation**:
  - HandlerFactory tạo UserHandler và các handlers khác
  - DatabaseFactory chọn PostgreSQL hoặc MongoDB dựa trên config

### 3. **Strategy Pattern**
- **File**: `usecase/user/password_strategy.go`
- **Mục đích**: Cho phép thay đổi thuật toán hash password
- **Implementation**: Interface `PasswordHasher` với các implementations: BcryptHasher, SHA256Hasher

### 4. **Observer Pattern**
- **File**: `infrastructure/observer/event.go`
- **Mục đích**: Thông báo các events (user created, updated, deleted) đến các observers
- **Implementation**: Subject-Observer pattern với UserEventLogger và UserEventNotifier

### 5. **Functional Options Pattern**
- **File**: `usecase/user/user_usecase.go`
- **Mục đích**: Cấu hình linh hoạt cho UserUseCase
- **Implementation**: 
  - `WithEmailValidation(bool)`
  - `WithPasswordValidation(bool)`
  - `WithPasswordLength(min, max)`

## 🚀 Cài đặt và Chạy

### Prerequisites
- Go 1.21+
- PostgreSQL 13+ **hoặc** MongoDB 7+ (tùy chọn)
- Docker & Docker Compose (khuyến nghị)

### 1. Clone và cài đặt dependencies
```bash
cd booking
go mod download
```

### 2. Khởi động Database với Docker
```bash
# Khởi động cả PostgreSQL và MongoDB
docker-compose up -d

# Hoặc chỉ PostgreSQL
docker-compose up -d postgres

# Hoặc chỉ MongoDB
docker-compose up -d mongodb
```

### 3. Cấu hình Database
Tạo file `.env` từ `.env.example`:
```bash
cp .env.example .env
```

Chỉnh sửa `.env` để chọn database type:

**Sử dụng PostgreSQL:**
```env
DB_TYPE=postgres
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=booking_db
DB_SSLMODE=disable
```

**Sử dụng MongoDB:**
```env
DB_TYPE=mongodb
MONGO_URI=mongodb://localhost:27017
MONGO_DB_NAME=booking_db
MONGO_TIMEOUT=10
```

### 4. Chạy Application
```bash
go run cmd/api/main.go
# hoặc
make run
```

Server sẽ chạy tại: `http://localhost:8080`

## 📚 API Endpoints

### Health Check
```
GET /health
```

### User CRUD Operations

#### Create User
```
POST /api/v1/users
Content-Type: application/json

{
  "email": "user@example.com",
  "username": "johndoe",
  "password": "securepassword123",
  "full_name": "John Doe",
  "phone": "+1234567890"
}
```

#### List Users
```
GET /api/v1/users?limit=10&offset=0&is_active=true
```

#### Get User by ID
```
GET /api/v1/users/:id
```

#### Update User
```
PUT /api/v1/users/:id
Content-Type: application/json

{
  "full_name": "John Updated",
  "phone": "+9876543210"
}
```

#### Delete User
```
DELETE /api/v1/users/:id
```

## 🧪 Testing với cURL

### Create User
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "username": "testuser",
    "password": "password123",
    "full_name": "Test User"
  }'
```

### Get All Users
```bash
curl http://localhost:8080/api/v1/users
```

### Get User by ID
```bash
curl http://localhost:8080/api/v1/users/1
```

## 🔧 Cấu trúc Code

### Domain Layer
- **Entity**: `domain/entity/user.go` - User entity với GORM tags
- **Repository Interface**: `domain/repository/user_repository.go` - Định nghĩa contract cho data access

### Use Case Layer
- **User UseCase**: `usecase/user/user_usecase.go` - Business logic
- **Password Strategy**: `usecase/user/password_strategy.go` - Password hashing strategies
- **Validation**: `usecase/user/validation.go` - Input validation

### Infrastructure Layer
- **Database**: `infrastructure/database/postgres.go` - Singleton database connection
- **Repository Implementation**: `infrastructure/database/user_repository_impl.go`
- **Observer**: `infrastructure/observer/event.go` - Event system

### Delivery Layer
- **Handlers**: `delivery/http/handler/user_handler.go` - HTTP request handlers
- **Router**: `delivery/http/router.go` - Route configuration
- **Middleware**: `delivery/http/middleware/` - CORS, Logger

## 📝 Design Patterns trong Code

Khi tạo user mới, bạn sẽ thấy các patterns hoạt động:

1. **Singleton**: Database connection được tái sử dụng
2. **Factory**: HandlerFactory tạo UserHandler
3. **Strategy**: BcryptHasher được sử dụng để hash password
4. **Observer**: UserEventLogger và UserEventNotifier được thông báo
5. **Functional Options**: UseCase được cấu hình với validation options

## 🎓 Học từ Code

Mỗi file đều có comments giải thích pattern được sử dụng. Đọc code theo thứ tự:

1. `domain/` - Hiểu entities và interfaces
2. `infrastructure/` - Xem implementations và patterns
3. `usecase/` - Học business logic và functional options
4. `delivery/` - Hiểu HTTP layer và routing
5. `cmd/api/main.go` - Xem cách tất cả được wire together

## 🔐 Security Notes

- Passwords được hash với bcrypt (cost factor 10)
- Password không được expose trong JSON responses
- Input validation được thực hiện ở use case layer
- CORS middleware được cấu hình

## 📦 Dependencies

- **gin-gonic/gin**: Web framework
- **gorm.io/gorm**: ORM
- **gorm.io/driver/postgres**: PostgreSQL driver
- **joho/godotenv**: Environment variables
- **golang.org/x/crypto**: Bcrypt hashing

