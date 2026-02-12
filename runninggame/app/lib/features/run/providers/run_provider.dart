import 'package:riverpod_annotation/riverpod_annotation.dart';
import '../../../core/models/run_model.dart';
import '../data/run_repository.dart';

part 'run_provider.g.dart';

@riverpod
class Run extends _$Run {
  @override
  RunModel? build() {
    return null;
  }

  Future<void> submitRun({
    required int distance,
    required int duration,
    required String startedAt,
    required String endedAt,
  }) async {
    final runRepository = ref.read(runRepositoryProvider);
    final result = await runRepository.submitRun(
      distance: distance,
      duration: duration,
      startedAt: startedAt,
      endedAt: endedAt,
    );
    state = result;
  }
}

@riverpod
Future<List<RunModel>> userRuns(UserRunsRef ref) async {
  final runRepository = ref.watch(runRepositoryProvider);
  return await runRepository.getUserRuns();
}

@riverpod
Future<RunStatsModel> userStats(UserStatsRef ref) async {
  final runRepository = ref.watch(runRepositoryProvider);
  return await runRepository.getUserStats();
}

