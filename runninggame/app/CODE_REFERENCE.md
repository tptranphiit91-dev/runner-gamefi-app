# Hype Stride - Code Reference Guide

Quick reference for key implementations in the Hype Stride Flutter app.

## 1. Main Entry Point

**File:** `lib/main.dart`

```dart
void main() {
  runApp(const ProviderScope(child: HypeStrideApp()));
}
```

## 2. Dio Client with JWT Interceptor

**File:** `lib/core/network/dio_client.dart`

```dart
@riverpod
Dio dio(DioRef ref) {
  final dio = Dio(BaseOptions(
    baseUrl: ApiConstants.baseUrl,
    connectTimeout: ApiConstants.connectTimeout,
  ));

  dio.interceptors.add(InterceptorsWrapper(
    onRequest: (options, handler) async {
      final prefs = await SharedPreferences.getInstance();
      final token = prefs.getString(StorageKeys.jwtToken);
      if (token != null) {
        options.headers['Authorization'] = 'Bearer $token';
      }
      return handler.next(options);
    },
  ));

  return dio;
}
```

## 3. GoRouter Configuration

**File:** `lib/core/router/app_router.dart`

```dart
@riverpod
GoRouter goRouter(GoRouterRef ref) {
  final authState = ref.watch(authStateProvider);

  return GoRouter(
    initialLocation: '/login',
    redirect: (context, state) {
      final isLoggedIn = authState.asData?.value ?? false;
      final isLoggingIn = state.matchedLocation == '/login';

      if (!isLoggedIn && !isLoggingIn) return '/login';
      if (isLoggedIn && isLoggingIn) return '/home';
      return null;
    },
    routes: [
      GoRoute(path: '/login', builder: (context, state) => const LoginScreen()),
      GoRoute(path: '/home', builder: (context, state) => const HomeScreen()),
      GoRoute(path: '/run', builder: (context, state) => const RunScreen()),
      GoRoute(path: '/shop', builder: (context, state) => const ShopScreen()),
    ],
  );
}
```

## 4. Login Implementation

**File:** `lib/features/auth/data/auth_repository.dart`

```dart
Future<Map<String, dynamic>> login({
  required String email,
  required String password,
}) async {
  final response = await _dio.post(
    ApiConstants.login,
    data: {'email': email, 'password': password},
  );

  final token = response.data['token'] as String;
  final user = UserModel.fromJson(response.data['user']);

  final prefs = await SharedPreferences.getInstance();
  await prefs.setString(StorageKeys.jwtToken, token);
  
  return {'token': token, 'user': user};
}
```

## 5. GPS Tracking (RunScreen)

**File:** `lib/features/run/presentation/run_screen.dart`

```dart
// Start GPS tracking
const LocationSettings locationSettings = LocationSettings(
  accuracy: LocationAccuracy.high,
  distanceFilter: 10,
);

_positionStream = Geolocator.getPositionStream(
  locationSettings: locationSettings,
).listen((Position position) {
  if (_lastPosition != null) {
    double distance = Geolocator.distanceBetween(
      _lastPosition!.latitude,
      _lastPosition!.longitude,
      position.latitude,
      position.longitude,
    );
    setState(() => _totalDistance += distance);
  }
  _lastPosition = position;
});
```

## 6. Submit Run to Backend

**File:** `lib/features/run/data/run_repository.dart`

```dart
Future<RunModel> submitRun({
  required int distance,
  required int duration,
  required String startedAt,
  required String endedAt,
}) async {
  final response = await _dio.post(
    ApiConstants.runs,
    data: {
      'distance': distance,
      'duration': duration,
      'started_at': startedAt,
      'ended_at': endedAt,
    },
  );

  return RunModel.fromJson(response.data['data']);
}
```

## 7. Shop Items Grid

**File:** `lib/features/shop/presentation/shop_screen.dart`

```dart
final itemsAsync = ref.watch(shopItemsProvider);

return itemsAsync.when(
  data: (items) => GridView.builder(
    gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
      crossAxisCount: 2,
      childAspectRatio: 0.75,
    ),
    itemCount: items.length,
    itemBuilder: (context, index) => _ItemCard(item: items[index]),
  ),
  loading: () => const CircularProgressIndicator(),
  error: (error, stack) => Text('Error: $error'),
);
```

## 8. Buy Item

**File:** `lib/features/shop/data/shop_repository.dart`

```dart
Future<void> buyItem(String itemId) async {
  await _dio.post(
    ApiConstants.buyItem,
    data: {'item_id': itemId},
  );
}
```

## 9. Freezed Model Example

**File:** `lib/core/models/user_model.dart`

```dart
@freezed
class UserModel with _$UserModel {
  const factory UserModel({
    required String id,
    required String username,
    required String email,
    @JsonKey(name: 'wallet_balance') required double walletBalance,
  }) = _UserModel;

  factory UserModel.fromJson(Map<String, dynamic> json) =>
      _$UserModelFromJson(json);
}
```

## 10. Riverpod Provider Example

**File:** `lib/features/auth/providers/login_provider.dart`

```dart
@riverpod
class Login extends _$Login {
  @override
  UserModel? build() => null;

  Future<void> login({
    required String email,
    required String password,
  }) async {
    final authRepository = ref.read(authRepositoryProvider);
    final result = await authRepository.login(
      email: email,
      password: password,
    );
    state = result['user'] as UserModel;
    ref.invalidate(authStateProvider);
  }
}
```

## Common Commands

```bash
# Setup project
make setup

# Run code generation
flutter pub run build_runner build --delete-conflicting-outputs

# Watch for changes
flutter pub run build_runner watch

# Run app
flutter run

# Clean and rebuild
make rebuild
```

## API Endpoints Used

```dart
// lib/core/constants/api_constants.dart
static const String baseUrl = 'http://localhost:8080';
static const String login = '/api/v1/auth/login';
static const String runs = '/api/v1/runs';
static const String items = '/api/v1/items';
static const String buyItem = '/api/v1/items/buy';
```

## Navigation

```dart
// Navigate to different screens
context.go('/home');
context.go('/run');
context.go('/shop');
context.go('/login');
```

## Error Handling

```dart
try {
  await someAsyncOperation();
} catch (e) {
  ScaffoldMessenger.of(context).showSnackBar(
    SnackBar(
      content: Text(e.toString().replaceAll('Exception: ', '')),
      backgroundColor: Colors.red,
    ),
  );
}
```

