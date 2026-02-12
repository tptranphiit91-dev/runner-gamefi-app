# 🍃 MongoDB Integration Guide

## Tổng quan

Booking Service hiện hỗ trợ **cả PostgreSQL và MongoDB** databases. Bạn có thể chọn database nào sử dụng thông qua biến môi trường `DB_TYPE`.

## 🏗️ Kiến trúc

### Database Abstraction Layer

```
Domain Layer (Interface)
    ↓
repository.UserRepository (interface)
    ↓
    ├── PostgreSQL Implementation (user_repository_impl.go)
    └── MongoDB Implementation (user_repository_mongo.go)
```

### Factory Pattern

DatabaseFactory tự động chọn implementation phù hợp:

```go
DB_TYPE=postgres → PostgreSQL Repository
DB_TYPE=mongodb  → MongoDB Repository
```

## 📁 Files mới

### 1. `infrastructure/database/mongodb.go`
- MongoDB connection với Singleton pattern
- Thread-safe initialization với `sync.Once`
- Connection pooling và health check

### 2. `infrastructure/database/user_repository_mongo.go`
- Implementation của `repository.UserRepository` cho MongoDB
- Tất cả CRUD operations
- Tích hợp Observer pattern cho events

### 3. `infrastructure/database/database_factory.go`
- Factory Pattern để chọn database type
- Tạo repository phù hợp dựa trên config
- Quản lý lifecycle của database connections

## 🚀 Cách sử dụng

### 1. Khởi động MongoDB

**Với Docker:**
```bash
docker-compose up -d mongodb
```

**Thủ công (macOS):**
```bash
brew tap mongodb/brew
brew install mongodb-community@7
brew services start mongodb-community@7
```

### 2. Cấu hình .env

```env
# Chọn MongoDB
DB_TYPE=mongodb

# MongoDB Configuration
MONGO_URI=mongodb://localhost:27017
MONGO_DB_NAME=booking_db
MONGO_TIMEOUT=10
```

### 3. Chạy Application

```bash
go run cmd/api/main.go
```

Bạn sẽ thấy:
```
🔧 Database Type: mongodb
✅ Observers attached
✅ Database connected successfully (mongodb)
```

## 🔄 Chuyển đổi giữa Databases

Chỉ cần thay đổi `DB_TYPE` trong `.env`:

**PostgreSQL:**
```env
DB_TYPE=postgres
```

**MongoDB:**
```env
DB_TYPE=mongodb
```

Không cần thay đổi code! Factory Pattern tự động xử lý.

## 📊 So sánh

| Feature | PostgreSQL | MongoDB |
|---------|-----------|---------|
| **Type** | Relational | Document |
| **Schema** | Strict | Flexible |
| **Transactions** | Full ACID | ACID (4.0+) |
| **Queries** | SQL | BSON/MQL |
| **Indexes** | B-tree | B-tree, Text, Geo |
| **Use Case** | Structured data | Semi-structured |

## 🔍 Implementation Details

### MongoDB User Document

```go
type MongoUser struct {
    ID        primitive.ObjectID `bson:"_id,omitempty"`
    Email     string             `bson:"email"`
    Username  string             `bson:"username"`
    Password  string             `bson:"password"`
    FullName  string             `bson:"full_name"`
    Phone     string             `bson:"phone"`
    IsActive  bool               `bson:"is_active"`
    CreatedAt time.Time          `bson:"created_at"`
    UpdatedAt time.Time          `bson:"updated_at"`
}
```

### Indexes

Tự động tạo unique indexes cho:
- `email` (unique)
- `username` (unique)

### ID Mapping

MongoDB sử dụng `ObjectID` (12 bytes) trong khi domain entity sử dụng `uint`.
Conversion được xử lý tự động trong repository layer.

## ✅ Features hỗ trợ

Tất cả operations đều hoạt động với cả 2 databases:

- ✅ Create User
- ✅ Get User by ID
- ✅ Get User by Email
- ✅ Get User by Username
- ✅ List Users (với filters)
- ✅ Update User
- ✅ Delete User
- ✅ Count Users
- ✅ Observer Events (UserCreated, UserUpdated, UserDeleted)

## 🧪 Testing

Test với PostgreSQL:
```bash
DB_TYPE=postgres go run cmd/api/main.go
./test_api.sh
```

Test với MongoDB:
```bash
DB_TYPE=mongodb go run cmd/api/main.go
./test_api.sh
```

## 🎯 Design Patterns sử dụng

1. **Singleton Pattern** - MongoDB connection
2. **Factory Pattern** - Database selection
3. **Repository Pattern** - Data access abstraction
4. **Observer Pattern** - Event notifications

## 📝 Notes

### Limitations

- ID conversion từ ObjectID sang uint là simplified approach
- Production systems nên sử dụng ObjectID trực tiếp hoặc maintain ID mapping table
- Một số advanced MongoDB features chưa được sử dụng (aggregation pipeline, etc.)

### Best Practices

- Luôn set `DB_TYPE` trong `.env`
- Sử dụng connection pooling (đã được config sẵn)
- Monitor database connections
- Backup data thường xuyên

## 🔗 Related Documentation

- [README.md](README.md) - Main documentation
- [DESIGN_PATTERNS.md](DESIGN_PATTERNS.md) - Design patterns details
- [QUICKSTART.md](QUICKSTART.md) - Quick start guide

---

**Tạo bởi**: Booking Service Team  
**Ngày**: 2026-01-28  
**Version**: 1.0.0

