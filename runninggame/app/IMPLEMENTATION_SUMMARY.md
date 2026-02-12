# Hype Stride - Implementation Summary

## Overview

A complete Flutter application for the Hype Stride gamified running app, built with modern architecture and best practices.

## Tech Stack Implementation

### ✅ State Management: Riverpod with Code Generation
- All providers use `@riverpod` annotation for type-safe code generation
- Providers located in feature-specific `providers/` folders
- Auth state management with automatic router integration

### ✅ Networking: Dio
- Centralized `DioClient` with automatic JWT token injection
- Interceptors for authentication and error handling
- Repository pattern for API calls

### ✅ Routing: GoRouter
- Declarative routing with path-based navigation
- Authentication-aware redirects
- Bottom navigation integration

### ✅ Local Storage: SharedPreferences
- JWT token persistence
- User data caching
- Automatic token attachment to API requests

### ✅ GPS Tracking: Geolocator
- Real-time position tracking
- Distance calculation between coordinates
- Permission handling for iOS and Android

## Project Structure

```
lib/
├── core/
│   ├── constants/
│   │   ├── api_constants.dart          # API endpoints and configuration
│   │   └── storage_keys.dart           # SharedPreferences keys
│   ├── models/
│   │   ├── user_model.dart             # User data model (Freezed)
│   │   ├── run_model.dart              # Run and stats models (Freezed)
│   │   └── item_model.dart             # Shop item model (Freezed)
│   ├── network/
│   │   ├── dio_client.dart             # Dio instance with JWT interceptor
│   │   └── api_response.dart           # Generic API response wrapper
│   ├── providers/
│   │   └── auth_state_provider.dart    # Authentication state
│   └── router/
│       └── app_router.dart             # GoRouter configuration
├── features/
│   ├── auth/
│   │   ├── data/
│   │   │   └── auth_repository.dart    # Login, register, logout
│   │   ├── presentation/
│   │   │   └── login_screen.dart       # Email/password login UI
│   │   └── providers/
│   │       └── login_provider.dart     # Login state management
│   ├── home/
│   │   └── presentation/
│   │       └── home_screen.dart        # The Crib - user dashboard
│   ├── run/
│   │   ├── data/
│   │   │   └── run_repository.dart     # Submit runs, fetch stats
│   │   ├── presentation/
│   │   │   └── run_screen.dart         # GPS tracking UI
│   │   └── providers/
│   │       └── run_provider.dart       # Run state management
│   └── shop/
│       ├── data/
│       │   └── shop_repository.dart    # Fetch items, buy, equip
│       ├── presentation/
│       │   └── shop_screen.dart        # Shop grid view
│       └── providers/
│           └── shop_provider.dart      # Shop state management
└── main.dart                           # App entry point
```

## Key Features Implementation

### 1. Authentication Flow

**LoginScreen** (`features/auth/presentation/login_screen.dart`)
- Email/password form with validation
- Loading state during authentication
- Error handling with SnackBar
- Automatic navigation to home on success

**AuthRepository** (`features/auth/data/auth_repository.dart`)
- Login: POST to `/api/v1/auth/login`
- Saves JWT token to SharedPreferences
- Provides `isLoggedIn()` check for router

**Router Integration** (`core/router/app_router.dart`)
- Watches `authStateProvider`
- Redirects unauthenticated users to `/login`
- Redirects authenticated users away from `/login`

### 2. Home Screen - "The Crib"

**HomeScreen** (`features/home/presentation/home_screen.dart`)
- Displays user avatar (placeholder)
- Shows wallet balance in AppBar
- Room customization area (placeholder for future items)
- Bottom navigation bar
- Logout functionality

### 3. Run Tracking

**RunScreen** (`features/run/presentation/run_screen.dart`)

**GPS Tracking Implementation:**
```dart
// Permission checking
LocationPermission permission = await Geolocator.checkPermission();

// Position stream with high accuracy
const LocationSettings locationSettings = LocationSettings(
  accuracy: LocationAccuracy.high,
  distanceFilter: 10, // Update every 10 meters
);

// Distance calculation
double distance = Geolocator.distanceBetween(
  lastPosition.latitude,
  lastPosition.longitude,
  currentPosition.latitude,
  currentPosition.longitude,
);
```

**Features:**
- Real-time timer (HH:MM:SS format)
- Distance tracking in kilometers
- Start/Stop run functionality
- Automatic submission to backend on stop
- ISO 8601 timestamp formatting

**RunRepository** (`features/run/data/run_repository.dart`)
- Submit run: POST to `/api/v1/runs`
- Payload: `{distance, duration, started_at, ended_at}`
- Returns earned coins from backend

### 4. Shop

**ShopScreen** (`features/shop/presentation/shop_screen.dart`)
- GridView of available items
- Item cards with image, name, type, price
- Buy confirmation dialog
- Error handling for insufficient balance
- Auto-refresh after purchase

**ShopRepository** (`features/shop/data/shop_repository.dart`)
- Get items: GET `/api/v1/items`
- Buy item: POST `/api/v1/items/buy`
- Equip/unequip functionality

## Network Layer

### DioClient with JWT Interceptor

**Automatic Token Injection:**
```dart
dio.interceptors.add(
  InterceptorsWrapper(
    onRequest: (options, handler) async {
      final prefs = await SharedPreferences.getInstance();
      final token = prefs.getString(StorageKeys.jwtToken);
      
      if (token != null) {
        options.headers['Authorization'] = 'Bearer $token';
      }
      return handler.next(options);
    },
  ),
);
```

**401 Handling:**
- Automatically clears token on 401 response
- Removes user data from SharedPreferences
- User redirected to login by router

## Material 3 Design

- Color scheme based on blue seed color
- Rounded corners on cards and buttons
- Elevation and shadows for depth
- Consistent spacing and padding
- Bottom navigation for main screens

## Code Generation

### Required Generators:
1. **Riverpod Generator** - `@riverpod` annotations
2. **Freezed** - `@freezed` for immutable models
3. **JSON Serializable** - `@JsonSerializable` for API models

### Build Command:
```bash
flutter pub run build_runner build --delete-conflicting-outputs
```

## Platform Configuration

### iOS (Info.plist)
```xml
<key>NSLocationWhenInUseUsageDescription</key>
<string>We need your location to track your runs</string>
```

### Android (AndroidManifest.xml)
```xml
<uses-permission android:name="android.permission.ACCESS_FINE_LOCATION" />
<uses-permission android:name="android.permission.ACCESS_COARSE_LOCATION" />
```

## API Integration

All endpoints from backend API documentation are implemented:

| Endpoint | Method | Implementation |
|----------|--------|----------------|
| `/api/v1/auth/login` | POST | AuthRepository.login() |
| `/api/v1/auth/register` | POST | AuthRepository.register() |
| `/api/v1/runs` | POST | RunRepository.submitRun() |
| `/api/v1/runs` | GET | RunRepository.getUserRuns() |
| `/api/v1/runs/stats` | GET | RunRepository.getUserStats() |
| `/api/v1/items` | GET | ShopRepository.getAvailableItems() |
| `/api/v1/items/buy` | POST | ShopRepository.buyItem() |
| `/api/v1/items/equip` | POST | ShopRepository.equipItem() |
| `/api/v1/items/unequip` | POST | ShopRepository.unequipItem() |

## Next Steps

1. Run `make setup` to install dependencies and generate code
2. Update `api_constants.dart` with your backend URL
3. Run the backend server
4. Run the app with `flutter run`
5. Test the complete flow: Login → Run → Shop

## Future Enhancements

- [ ] User registration screen
- [ ] Run history list
- [ ] Statistics dashboard
- [ ] Room customization with owned items
- [ ] Avatar customization
- [ ] Social features (leaderboards)
- [ ] Offline support with local database
- [ ] Push notifications

