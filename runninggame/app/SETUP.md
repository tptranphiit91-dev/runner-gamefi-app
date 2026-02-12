# Hype Stride - Setup Guide

## Quick Start

Follow these steps to get the Hype Stride Flutter app running:

### 1. Prerequisites

Ensure you have the following installed:
- Flutter SDK (>=3.0.0) - [Install Flutter](https://flutter.dev/docs/get-started/install)
- Dart SDK (comes with Flutter)
- Xcode (for iOS development) or Android Studio (for Android development)
- A code editor (VS Code or Android Studio recommended)

### 2. Install Dependencies

```bash
cd runninggame/app
flutter pub get
```

### 3. Generate Code

The app uses code generation for Riverpod providers, Freezed models, and JSON serialization:

```bash
flutter pub run build_runner build --delete-conflicting-outputs
```

**Note**: This step is crucial! The app won't compile without running code generation first.

### 4. Configure Backend URL

Update the backend API URL in `lib/core/constants/api_constants.dart`:

```dart
static const String baseUrl = 'http://YOUR_IP_ADDRESS:8080';
```

**Important**: 
- For iOS Simulator: Use your computer's local IP address (not `localhost`)
- For Android Emulator: Use `10.0.2.2` instead of `localhost`
- For Physical Devices: Use your computer's IP address on the same network

To find your IP address:
- macOS/Linux: Run `ifconfig | grep inet`
- Windows: Run `ipconfig`

### 5. Run the Backend

Make sure the backend API is running before starting the app:

```bash
cd ../backend
make run
```

See `../backend/README.md` for backend setup instructions.

### 6. Run the App

```bash
flutter run
```

Or select a device in your IDE and click Run.

## Platform-Specific Setup

### iOS Setup

1. Open `ios/Runner.xcworkspace` in Xcode
2. Select a development team in Signing & Capabilities
3. Location permissions are already configured in `Info.plist`

### Android Setup

1. Location permissions are already configured in `AndroidManifest.xml`
2. Make sure you have an Android emulator or device connected
3. Run `flutter devices` to verify

## Code Generation During Development

When developing and modifying files with annotations (`@riverpod`, `@freezed`, etc.), use:

```bash
flutter pub run build_runner watch
```

This will automatically regenerate code when you save files.

## Troubleshooting

### "No such file" errors for `.g.dart` or `.freezed.dart` files

Run code generation:
```bash
flutter pub run build_runner build --delete-conflicting-outputs
```

### Location permissions not working

**iOS**: Check `ios/Runner/Info.plist` contains location usage descriptions

**Android**: Check `android/app/src/main/AndroidManifest.xml` contains location permissions

### Cannot connect to backend

1. Verify backend is running: `curl http://localhost:8080/health`
2. Check the IP address in `api_constants.dart`
3. For iOS Simulator, use your computer's IP, not `localhost`
4. For Android Emulator, use `10.0.2.2` instead of `localhost`

### Build errors after pulling changes

```bash
flutter clean
flutter pub get
flutter pub run build_runner build --delete-conflicting-outputs
```

## Testing

### Test Login Credentials

After registering a user via the backend, use those credentials to login.

Example registration (using the backend API):
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testrunner",
    "email": "test@example.com",
    "password": "password123"
  }'
```

Then login in the app with:
- Email: `test@example.com`
- Password: `password123`

## Project Structure

```
lib/
├── core/               # Core functionality (network, routing, models)
├── features/           # Feature modules (auth, home, run, shop)
└── main.dart          # App entry point
```

## Next Steps

1. Run the app and login
2. Navigate to the Run screen and start tracking
3. Visit the Shop to see available items
4. Check your wallet balance on the Home screen

## Additional Resources

- [Flutter Documentation](https://flutter.dev/docs)
- [Riverpod Documentation](https://riverpod.dev)
- [GoRouter Documentation](https://pub.dev/packages/go_router)
- [Geolocator Documentation](https://pub.dev/packages/geolocator)

