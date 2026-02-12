import 'package:riverpod_annotation/riverpod_annotation.dart';
import '../../../core/models/user_model.dart';
import '../data/auth_repository.dart';

part 'login_provider.g.dart';

@riverpod
class Login extends _$Login {
  @override
  UserModel? build() {
    return null;
  }

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
    
    // Invalidate auth state to trigger router redirect
    ref.invalidate(authStateProvider);
  }

  Future<void> logout() async {
    final authRepository = ref.read(authRepositoryProvider);
    await authRepository.logout();
    state = null;
    
    // Invalidate auth state to trigger router redirect
    ref.invalidate(authStateProvider);
  }
}

// Import for authStateProvider
import '../../../core/providers/auth_state_provider.dart';

