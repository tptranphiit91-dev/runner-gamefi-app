# Hype Stride Flutter App - Project Overview

## 🎯 Project Summary

A complete, production-ready Flutter application for the Hype Stride gamified running app. Built with modern architecture patterns, type-safe state management, and Material 3 design.

## 📋 What's Included

### ✅ Complete Feature Set
- **Authentication**: Login with email/password, JWT token management
- **Home Screen**: User dashboard with avatar and wallet balance
- **Run Tracking**: Real-time GPS tracking with distance and time
- **Shop**: Browse and purchase items with coins earned from runs

### ✅ Tech Stack (As Requested)
- **Framework**: Flutter (Latest version)
- **State Management**: Riverpod with Code Generation annotations
- **Networking**: Dio with automatic JWT token injection
- **Routing**: GoRouter with authentication-aware redirects
- **Local Storage**: SharedPreferences for JWT token persistence
- **GPS Tracking**: Geolocator package for position tracking
- **UI**: Material 3 Design System

### ✅ Architecture
- **Feature-based folder structure**
- **Repository pattern** for data layer
- **Provider pattern** for state management
- **Freezed models** for immutability
- **JSON serialization** for API models

## 📁 Project Structure

```
runninggame/app/
├── lib/
│   ├── core/                          # Core functionality
│   │   ├── constants/                 # API URLs, storage keys
│   │   ├── models/                    # Shared data models
│   │   ├── network/                   # Dio client, interceptors
│   │   ├── providers/                 # Global providers
│   │   └── router/                    # GoRouter configuration
│   ├── features/                      # Feature modules
│   │   ├── auth/                      # Authentication
│   │   │   ├── data/                  # AuthRepository
│   │   │   ├── presentation/          # LoginScreen
│   │   │   └── providers/             # Login state
│   │   ├── home/                      # Home/Dashboard
│   │   │   └── presentation/          # HomeScreen
│   │   ├── run/                       # Run tracking
│   │   │   ├── data/                  # RunRepository
│   │   │   ├── presentation/          # RunScreen with GPS
│   │   │   └── providers/             # Run state
│   │   └── shop/                      # Shop/Store
│   │       ├── data/                  # ShopRepository
│   │       ├── presentation/          # ShopScreen
│   │       └── providers/             # Shop state
│   └── main.dart                      # App entry point
├── android/                           # Android configuration
├── ios/                               # iOS configuration
├── pubspec.yaml                       # Dependencies
├── build.yaml                         # Code generation config
├── Makefile                           # Build commands
├── README.md                          # Main documentation
├── SETUP.md                           # Setup instructions
├── IMPLEMENTATION_SUMMARY.md          # Technical details
└── CODE_REFERENCE.md                  # Code snippets

Total Files Created: 30+
```

## 🚀 Quick Start

```bash
# 1. Navigate to app directory
cd runninggame/app

# 2. Setup (install deps + generate code)
make setup

# 3. Update backend URL in lib/core/constants/api_constants.dart
# Change baseUrl to your backend IP address

# 4. Run the app
flutter run
```

## 🔑 Key Features Implementation

### 1. Authentication Flow
- Login screen with form validation
- JWT token saved to SharedPreferences
- Automatic token injection in all API requests
- Router redirects based on auth state

### 2. GPS Run Tracking
- Real-time position tracking with Geolocator
- Distance calculation between GPS coordinates
- Timer with HH:MM:SS format
- Submit run data to backend on completion
- Earn coins based on distance

### 3. Shop System
- Grid view of available items
- Item cards with image, name, price
- Buy confirmation dialog
- Error handling for insufficient balance
- Auto-refresh after purchase

### 4. Navigation
- Bottom navigation bar on all main screens
- GoRouter for declarative routing
- Deep linking support ready
- Authentication-aware redirects

## 📱 Screens

| Route | Screen | Description |
|-------|--------|-------------|
| `/login` | LoginScreen | Email/password authentication |
| `/home` | HomeScreen | User dashboard "The Crib" |
| `/run` | RunScreen | GPS tracking with timer |
| `/shop` | ShopScreen | Browse and buy items |

## 🔌 API Integration

All backend endpoints are fully integrated:

- ✅ POST `/api/v1/auth/login` - User login
- ✅ POST `/api/v1/auth/register` - User registration
- ✅ POST `/api/v1/runs` - Submit run data
- ✅ GET `/api/v1/runs` - Get user runs
- ✅ GET `/api/v1/runs/stats` - Get user statistics
- ✅ GET `/api/v1/items` - Get shop items
- ✅ POST `/api/v1/items/buy` - Purchase item
- ✅ POST `/api/v1/items/equip` - Equip item
- ✅ POST `/api/v1/items/unequip` - Unequip item

## 🛠️ Development Tools

### Makefile Commands
```bash
make setup          # Initial setup
make clean          # Clean build artifacts
make build-runner   # Generate code once
make watch          # Watch mode for code gen
make run            # Run the app
make test           # Run tests
make format         # Format code
make lint           # Run linter
```

### Code Generation
Required for Riverpod, Freezed, and JSON serialization:
```bash
flutter pub run build_runner build --delete-conflicting-outputs
```

## 📦 Dependencies

### Production
- `flutter_riverpod: ^2.5.1` - State management
- `riverpod_annotation: ^2.3.5` - Code generation
- `dio: ^5.4.0` - HTTP client
- `go_router: ^13.0.0` - Routing
- `shared_preferences: ^2.2.2` - Local storage
- `geolocator: ^11.0.0` - GPS tracking
- `freezed_annotation: ^2.4.1` - Immutable models
- `json_annotation: ^4.8.1` - JSON serialization

### Development
- `build_runner: ^2.4.8` - Code generation
- `riverpod_generator: ^2.3.11` - Riverpod codegen
- `freezed: ^2.4.6` - Model codegen
- `json_serializable: ^6.7.1` - JSON codegen
- `flutter_lints: ^3.0.0` - Linting

## 📱 Platform Support

### iOS
- ✅ Location permissions configured
- ✅ Info.plist setup complete
- ✅ Ready to run on simulator/device

### Android
- ✅ Location permissions configured
- ✅ AndroidManifest.xml setup complete
- ✅ Ready to run on emulator/device

## 📚 Documentation

- **README.md** - Main documentation and features
- **SETUP.md** - Detailed setup instructions
- **IMPLEMENTATION_SUMMARY.md** - Technical implementation details
- **CODE_REFERENCE.md** - Quick code snippets reference
- **PROJECT_OVERVIEW.md** - This file

## ✨ Code Quality

- ✅ Type-safe with Riverpod code generation
- ✅ Immutable models with Freezed
- ✅ Null-safety enabled
- ✅ Linting rules configured
- ✅ Material 3 design system
- ✅ Error handling throughout
- ✅ Loading states for async operations

## 🎨 UI/UX Features

- Material 3 design with blue color scheme
- Smooth animations and transitions
- Loading indicators for async operations
- Error messages with SnackBars
- Confirmation dialogs for important actions
- Bottom navigation for easy access
- Responsive layouts

## 🔐 Security

- JWT token stored securely in SharedPreferences
- Automatic token injection in API requests
- Token cleared on 401 responses
- Sensitive data not logged in production

## 🚦 Next Steps

1. **Run the backend**: `cd ../backend && make run`
2. **Setup the app**: `make setup`
3. **Update API URL**: Edit `lib/core/constants/api_constants.dart`
4. **Run the app**: `flutter run`
5. **Test the flow**: Login → Run → Shop

## 📝 Notes

- All code uses latest Flutter/Dart best practices
- Ready for production deployment
- Extensible architecture for future features
- Well-documented and maintainable code

---

**Built with ❤️ using Flutter and Riverpod**

