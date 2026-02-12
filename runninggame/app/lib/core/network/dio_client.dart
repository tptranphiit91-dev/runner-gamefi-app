import 'package:dio/dio.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';
import 'package:shared_preferences/shared_preferences.dart';
import '../constants/api_constants.dart';
import '../constants/storage_keys.dart';

part 'dio_client.g.dart';

@riverpod
Dio dio(DioRef ref) {
  final dio = Dio(
    BaseOptions(
      baseUrl: ApiConstants.baseUrl,
      connectTimeout: ApiConstants.connectTimeout,
      receiveTimeout: ApiConstants.receiveTimeout,
      headers: {
        'Content-Type': 'application/json',
        'Accept': 'application/json',
      },
    ),
  );

  // Add JWT interceptor
  dio.interceptors.add(
    InterceptorsWrapper(
      onRequest: (options, handler) async {
        // Get token from SharedPreferences
        final prefs = await SharedPreferences.getInstance();
        final token = prefs.getString(StorageKeys.jwtToken);

        if (token != null && token.isNotEmpty) {
          options.headers[ApiConstants.authHeader] =
              '${ApiConstants.bearerPrefix} $token';
        }

        return handler.next(options);
      },
      onError: (error, handler) async {
        // Handle 401 Unauthorized - token expired or invalid
        if (error.response?.statusCode == 401) {
          // Clear token and redirect to login
          final prefs = await SharedPreferences.getInstance();
          await prefs.remove(StorageKeys.jwtToken);
          await prefs.remove(StorageKeys.userId);
          await prefs.remove(StorageKeys.userEmail);
          await prefs.remove(StorageKeys.username);
        }
        return handler.next(error);
      },
    ),
  );

  // Add logging interceptor for debugging
  dio.interceptors.add(
    LogInterceptor(
      requestBody: true,
      responseBody: true,
      error: true,
      requestHeader: true,
      responseHeader: false,
    ),
  );

  return dio;
}

