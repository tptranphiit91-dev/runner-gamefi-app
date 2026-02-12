# 📚 Documentation Index

Chào mừng đến với Booking Service! Đây là hướng dẫn để navigate qua tất cả documentation.

---

## 🚀 Bắt đầu nhanh

### Bạn muốn gì?

#### "Tôi muốn chạy project ngay!"
👉 Đọc **[QUICKSTART.md](QUICKSTART.md)**
- 3 commands để chạy
- Test API ngay lập tức
- Troubleshooting

#### "Tôi muốn hiểu project này làm gì?"
👉 Đọc **[README.md](README.md)**
- Tổng quan về project
- Kiến trúc Clean Architecture
- API endpoints
- Cách cài đặt chi tiết

#### "Tôi muốn học Design Patterns!"
👉 Đọc **[DESIGN_PATTERNS.md](DESIGN_PATTERNS.md)**
- 5 Design Patterns chi tiết
- Code examples
- Benefits của từng pattern
- Cách kết hợp patterns

#### "Tôi muốn hiểu cấu trúc code?"
👉 Đọc **[PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md)**
- Cấu trúc thư mục
- File descriptions
- Request flow
- Component relationships

#### "Tôi muốn xem examples cụ thể?"
👉 Đọc **[EXAMPLES.md](EXAMPLES.md)**
- Thay đổi password strategy
- Thêm observers mới
- Thêm functional options
- API request examples
- Testing examples

#### "Tôi muốn tổng quan nhanh?"
👉 Đọc **[SUMMARY.md](SUMMARY.md)**
- Checklist những gì đã làm
- Statistics
- Key features
- Learning path

#### "Tôi muốn sử dụng MongoDB?"
👉 Đọc **[MONGODB_INTEGRATION.md](MONGODB_INTEGRATION.md)**
- MongoDB setup guide
- Cách chuyển đổi giữa PostgreSQL và MongoDB
- Database Factory pattern
- Implementation details

---

## 📖 Reading Path

### 🟢 Beginner (Mới bắt đầu)
1. **[QUICKSTART.md](QUICKSTART.md)** - Chạy project (5 phút)
2. **[README.md](README.md)** - Hiểu tổng quan (15 phút)
3. **[SUMMARY.md](SUMMARY.md)** - Xem tổng kết (10 phút)

**Thời gian**: ~30 phút

### 🟡 Intermediate (Đã biết cơ bản)
1. **[PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md)** - Hiểu cấu trúc (20 phút)
2. **[DESIGN_PATTERNS.md](DESIGN_PATTERNS.md)** - Học patterns (30 phút)
3. **[EXAMPLES.md](EXAMPLES.md)** - Xem examples (20 phút)
4. Đọc source code theo thứ tự trong PROJECT_STRUCTURE.md

**Thời gian**: ~2 giờ

### 🔴 Advanced (Muốn master)
1. Đọc tất cả documentation
2. Đọc toàn bộ source code
3. Chạy và debug từng API endpoint
4. Thử modify code và thêm features
5. Viết tests

**Thời gian**: ~1 ngày

---

## 📁 File Organization

### Documentation Files
```
📄 INDEX.md                  ← Bạn đang ở đây
📄 README.md                 ← Main documentation
📄 QUICKSTART.md             ← Quick start guide
📄 DESIGN_PATTERNS.md        ← Design patterns explained
📄 PROJECT_STRUCTURE.md      ← Project structure
📄 EXAMPLES.md               ← Code examples
📄 SUMMARY.md                ← Project summary
📄 MONGODB_INTEGRATION.md    ← MongoDB guide (NEW!)
```

### Configuration Files
```
⚙️ .env.example              ← Environment template
⚙️ docker-compose.yml        ← PostgreSQL & MongoDB containers
⚙️ Makefile                  ← Build commands
⚙️ go.mod                    ← Go dependencies
```

### Source Code
```
📂 cmd/api/                  ← Application entry point
📂 config/                   ← Configuration
📂 delivery/http/            ← HTTP layer
📂 usecase/user/             ← Business logic
📂 domain/                   ← Entities & interfaces
📂 infrastructure/           ← Database & external
```

### Tools
```
🔧 test_api.sh               ← API testing script
🔧 Makefile                  ← Build & run commands
```

---

## 🎯 Use Cases

### "Tôi muốn học Clean Architecture"
1. Đọc README.md (Architecture section)
2. Đọc PROJECT_STRUCTURE.md (Layer Dependencies)
3. Đọc code theo thứ tự: Domain → UseCase → Infrastructure → Delivery

### "Tôi muốn học Design Patterns"
1. Đọc DESIGN_PATTERNS.md
2. Đọc EXAMPLES.md
3. Xem code implementation trong từng file
4. Thử modify patterns

### "Tôi muốn build API tương tự"
1. Đọc QUICKSTART.md để chạy
2. Đọc PROJECT_STRUCTURE.md để hiểu cấu trúc
3. Copy structure và modify theo nhu cầu
4. Đọc EXAMPLES.md để biết cách extend

### "Tôi muốn contribute"
1. Đọc tất cả documentation
2. Chạy project và test
3. Đọc source code
4. Tìm areas để improve
5. Submit PR

---

## 🔍 Quick Reference

### Commands
```bash
make help          # Xem tất cả commands
make run           # Chạy service
make build         # Build binary
make test          # Run tests
docker-compose up  # Start PostgreSQL
./test_api.sh      # Test API
```

### API Endpoints
```
GET    /health              # Health check
POST   /api/v1/users        # Create user
GET    /api/v1/users        # List users
GET    /api/v1/users/:id    # Get user
PUT    /api/v1/users/:id    # Update user
DELETE /api/v1/users/:id    # Delete user
```

### Design Patterns Locations
```
Singleton          → infrastructure/database/postgres.go
                     infrastructure/database/mongodb.go
Factory            → delivery/http/handler/handler_factory.go
                     infrastructure/database/database_factory.go (NEW!)
Strategy           → usecase/user/password_strategy.go
Observer           → infrastructure/observer/event.go
Functional Options → usecase/user/user_usecase.go
```

---

## 💡 Tips

- 📖 Đọc documentation theo thứ tự phù hợp với level của bạn
- 🏃 Chạy code trước khi đọc để có context
- 🔍 Sử dụng search (Cmd/Ctrl + F) để tìm topics cụ thể
- 💻 Thử modify code để hiểu sâu hơn
- 📝 Ghi chú những gì bạn học được

---

## 🎓 Learning Goals

Sau khi hoàn thành, bạn sẽ hiểu:

✅ Clean Architecture principles  
✅ 5 Design Patterns quan trọng  
✅ SOLID Principles  
✅ Go best practices  
✅ RESTful API design  
✅ Database patterns  
✅ Testing strategies  
✅ Production-ready code structure  

---

## 📞 Need Help?

- Đọc QUICKSTART.md cho troubleshooting
- Xem EXAMPLES.md cho use cases cụ thể
- Đọc comments trong source code
- Check Makefile cho available commands

---

**Happy Learning! 🚀**

*Start with [QUICKSTART.md](QUICKSTART.md) if you're new!*

