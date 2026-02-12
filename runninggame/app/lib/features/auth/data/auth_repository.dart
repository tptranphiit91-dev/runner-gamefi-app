import 'package:dio/dio.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';
import 'package:shared_preferences/shared_preferences.dart';
import '../../../core/constants/api_constants.dart';
import '../../../core/constants/storage_keys.dart';
import '../../../core/models/user_model.dart';
import '../../../core/network/dio_client.dart';

part 'auth_repository.g.dart';

@riverpod
AuthRepository authRepository(AuthRepositoryRef ref) {
  return AuthRepository(ref.watch(dioProvider));
}

class AuthRepository {
  final Dio _dio;

  AuthRepository(this._dio);

  Future<Map<String, dynamic>> login({
    required String email,
    required String password,
  }) async {
    try {
      final response = await _dio.post(
        ApiConstants.login,
        data: {
          'email': email,
          'password': password,
        },
      );

      if (response.statusCode == 200) {
        final data = response.data;
        final token = data['token'] as String;
        final user = UserModel.fromJson(data['user']);

        // Save token and user data to SharedPreferences
        final prefs = await SharedPreferences.getInstance();
        await prefs.setString(StorageKeys.jwtToken, token);
        await prefs.setString(StorageKeys.userId, user.id);
        await prefs.setString(StorageKeys.userEmail, user.email);
        await prefs.setString(StorageKeys.username, user.username);

        return {
          'token': token,
          'user': user,
        };
      } else {
        throw Exception('Login failed');
      }
    } on DioException catch (e) {
      if (e.response?.statusCode == 401) {
        throw Exception('Invalid credentials');
      }
      throw Exception(e.response?.data['error'] ?? 'Login failed');
    }
  }

  Future<UserModel> register({
    required String username,
    required String email,
    required String password,
  }) async {
    try {
      final response = await _dio.post(
        ApiConstants.register,
        data: {
          'username': username,
          'email': email,
          'password': password,
        },
      );

      if (response.statusCode == 201) {
        return UserModel.fromJson(response.data['user']);
      } else {
        throw Exception('Registration failed');
      }
    } on DioException catch (e) {
      throw Exception(e.response?.data['error'] ?? 'Registration failed');
    }
  }

  Future<void> logout() async {
    final prefs = await SharedPreferences.getInstance();
    await prefs.remove(StorageKeys.jwtToken);
    await prefs.remove(StorageKeys.userId);
    await prefs.remove(StorageKeys.userEmail);
    await prefs.remove(StorageKeys.username);
  }

  Future<bool> isLoggedIn() async {
    final prefs = await SharedPreferences.getInstance();
    final token = prefs.getString(StorageKeys.jwtToken);
    return token != null && token.isNotEmpty;
  }

  Future<String?> getToken() async {
    final prefs = await SharedPreferences.getInstance();
    return prefs.getString(StorageKeys.jwtToken);
  }
}

