import 'package:riverpod_annotation/riverpod_annotation.dart';
import '../../features/auth/data/auth_repository.dart';

part 'auth_state_provider.g.dart';

@riverpod
Future<bool> authState(AuthStateRef ref) async {
  final authRepository = ref.watch(authRepositoryProvider);
  return await authRepository.isLoggedIn();
}

