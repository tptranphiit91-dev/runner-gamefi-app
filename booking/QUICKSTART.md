# 🚀 Quick Start Guide

Hướng dẫn nhanh để chạy Booking Service trong 5 phút.

## ⚡ Cách nhanh nhất (với Docker)

### 1. Start Database
```bash
# Khởi động cả PostgreSQL và MongoDB
docker-compose up -d

# Hoặc chỉ PostgreSQL
docker-compose up -d postgres

# Hoặc chỉ MongoDB
docker-compose up -d mongodb
```

### 2. Copy environment file
```bash
cp .env.example .env
```

**Chọn database type trong `.env`:**
- Để dùng PostgreSQL: `DB_TYPE=postgres` (mặc định)
- Để dùng MongoDB: `DB_TYPE=mongodb`

### 3. Run the service
```bash
make run
```

Hoặc:
```bash
go run cmd/api/main.go
```

### 4. Test API
```bash
# Health check
curl http://localhost:8080/health

# Create a user
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "username": "testuser",
    "password": "password123",
    "full_name": "Test User"
  }'

# Get all users
curl http://localhost:8080/api/v1/users
```

---

## 🗄️ Cài đặt Database thủ công

### 🐘 PostgreSQL

Nếu bạn không dùng Docker:

### macOS
```bash
brew install postgresql@15
brew services start postgresql@15
createdb booking_db
```

### Ubuntu/Debian
```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
sudo systemctl start postgresql
sudo -u postgres createdb booking_db
```

### Windows
1. Download PostgreSQL từ https://www.postgresql.org/download/windows/
2. Cài đặt và start service
3. Tạo database `booking_db`

### 🍃 MongoDB

Nếu bạn muốn dùng MongoDB:

#### macOS
```bash
brew tap mongodb/brew
brew install mongodb-community@7
brew services start mongodb-community@7
```

#### Ubuntu/Debian
```bash
wget -qO - https://www.mongodb.org/static/pgp/server-7.0.asc | sudo apt-key add -
echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu jammy/mongodb-org/7.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-7.0.list
sudo apt update
sudo apt install -y mongodb-org
sudo systemctl start mongod
```

#### Windows
1. Download MongoDB từ https://www.mongodb.com/try/download/community
2. Cài đặt và start service

---

## 📝 Cấu hình .env

Chỉnh sửa file `.env` nếu cần:

**Với PostgreSQL:**
```env
# Server
SERVER_HOST=0.0.0.0
SERVER_PORT=8080

# Database Type
DB_TYPE=postgres

# PostgreSQL
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=booking_db
DB_SSLMODE=disable
```

**Với MongoDB:**
```env
# Server
SERVER_HOST=0.0.0.0
SERVER_PORT=8080

# Database Type
DB_TYPE=mongodb

# MongoDB
MONGO_URI=mongodb://localhost:27017
MONGO_DB_NAME=booking_db
MONGO_TIMEOUT=10
```

---

## 🧪 Test toàn bộ API

Chạy script test tự động:

```bash
# Đảm bảo service đang chạy
make run

# Trong terminal khác, chạy test script
./test_api.sh
```

---

## 📦 Các lệnh Makefile hữu ích

```bash
make help          # Xem tất cả commands
make run           # Chạy service
make build         # Build binary
make test          # Chạy tests
make clean         # Xóa build artifacts
make docker-up     # Start Docker containers
make docker-down   # Stop Docker containers
```

---

## 🎯 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| POST | `/api/v1/users` | Create user |
| GET | `/api/v1/users` | List users |
| GET | `/api/v1/users/:id` | Get user by ID |
| PUT | `/api/v1/users/:id` | Update user |
| DELETE | `/api/v1/users/:id` | Delete user |

---

## 🔍 Xem Design Patterns

Đọc file `DESIGN_PATTERNS.md` để hiểu cách các patterns được implement:

- **Singleton**: Database connection
- **Factory**: Handler creation
- **Strategy**: Password hashing
- **Observer**: Event notifications
- **Functional Options**: UseCase configuration

---

## 🐛 Troubleshooting

### Lỗi: "connection refused"
- Kiểm tra PostgreSQL đang chạy: `pg_isready`
- Kiểm tra port 5432: `lsof -i :5432`

### Lỗi: "database does not exist"
```bash
createdb booking_db
```

### Lỗi: "port 8080 already in use"
Thay đổi `SERVER_PORT` trong `.env`:
```env
SERVER_PORT=8081
```

---

## 📚 Next Steps

1. Đọc `README.md` để hiểu kiến trúc
2. Đọc `DESIGN_PATTERNS.md` để học patterns
3. Xem code trong các folders:
   - `domain/` - Entities và interfaces
   - `usecase/` - Business logic
   - `infrastructure/` - Database và observers
   - `delivery/` - HTTP handlers

---

## 💡 Tips

- Xem logs khi tạo user để thấy Observer pattern hoạt động
- Thử thay đổi password strategy từ Bcrypt sang SHA256
- Thử thêm options mới cho UserUseCase
- Thử tạo handler mới với Factory pattern

Happy coding! 🎉

