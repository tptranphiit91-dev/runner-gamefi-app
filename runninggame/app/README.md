# Hype Stride

A gamified running app built with Flutter that rewards you for staying active.

## Features

- 🏃 **GPS Run Tracking**: Track your runs with real-time distance and time monitoring
- 💰 **Earn Coins**: Get rewarded with coins based on distance covered
- 🏠 **The Crib**: Customize your virtual room with items purchased from the shop
- 🛍️ **Shop**: Buy items to customize your avatar and room
- 🔐 **Authentication**: Secure login system with JWT tokens

## Tech Stack

- **Framework**: Flutter (Latest version)
- **State Management**: Riverpod with Code Generation
- **Networking**: Dio for REST API calls
- **Routing**: GoRouter
- **Local Storage**: SharedPreferences for JWT token storage
- **Location Services**: Geolocator for GPS tracking

## Project Structure

```
lib/
├── core/
│   ├── constants/
│   │   ├── api_constants.dart
│   │   └── storage_keys.dart
│   ├── models/
│   │   ├── user_model.dart
│   │   ├── run_model.dart
│   │   └── item_model.dart
│   ├── network/
│   │   ├── dio_client.dart
│   │   └── api_response.dart
│   ├── providers/
│   │   └── auth_state_provider.dart
│   └── router/
│       └── app_router.dart
├── features/
│   ├── auth/
│   │   ├── data/
│   │   │   └── auth_repository.dart
│   │   ├── presentation/
│   │   │   └── login_screen.dart
│   │   └── providers/
│   │       └── login_provider.dart
│   ├── home/
│   │   └── presentation/
│   │       └── home_screen.dart
│   ├── run/
│   │   ├── data/
│   │   │   └── run_repository.dart
│   │   ├── presentation/
│   │   │   └── run_screen.dart
│   │   └── providers/
│   │       └── run_provider.dart
│   └── shop/
│       ├── data/
│       │   └── shop_repository.dart
│       ├── presentation/
│       │   └── shop_screen.dart
│       └── providers/
│           └── shop_provider.dart
└── main.dart
```

## Getting Started

### Prerequisites

- Flutter SDK (>=3.0.0)
- Dart SDK
- iOS Simulator / Android Emulator or Physical Device
- Backend API running (see `../backend/README.md`)

### Installation

1. Navigate to the app directory:
```bash
cd runninggame/app
```

2. Install dependencies:
```bash
flutter pub get
```

3. Run code generation for Riverpod and Freezed:
```bash
flutter pub run build_runner build --delete-conflicting-outputs
```

4. Update the API base URL in `lib/core/constants/api_constants.dart`:
```dart
static const String baseUrl = 'http://YOUR_IP_ADDRESS:8080';
```

### Running the App

```bash
flutter run
```

### iOS Permissions

Add the following to `ios/Runner/Info.plist`:

```xml
<key>NSLocationWhenInUseUsageDescription</key>
<string>We need your location to track your runs</string>
<key>NSLocationAlwaysUsageDescription</key>
<string>We need your location to track your runs</string>
```

### Android Permissions

Permissions are already configured in the geolocator package. Make sure your `android/app/src/main/AndroidManifest.xml` includes:

```xml
<uses-permission android:name="android.permission.ACCESS_FINE_LOCATION" />
<uses-permission android:name="android.permission.ACCESS_COARSE_LOCATION" />
```

## API Endpoints

The app connects to the following endpoints:

- `POST /api/v1/auth/login` - User authentication
- `POST /api/v1/auth/register` - User registration
- `POST /api/v1/runs` - Submit run data
- `GET /api/v1/runs` - Get user runs
- `GET /api/v1/runs/stats` - Get user statistics
- `GET /api/v1/items` - Get shop items
- `POST /api/v1/items/buy` - Purchase item

## Development

### Code Generation

When you modify models or providers with annotations, run:

```bash
flutter pub run build_runner watch
```

This will automatically regenerate code when files change.

### Clean Build

If you encounter issues:

```bash
flutter clean
flutter pub get
flutter pub run build_runner build --delete-conflicting-outputs
```

## License

This project is part of the Hype Stride application suite.

