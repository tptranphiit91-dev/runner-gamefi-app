import 'package:dio/dio.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';
import '../../../core/constants/api_constants.dart';
import '../../../core/models/run_model.dart';
import '../../../core/network/dio_client.dart';

part 'run_repository.g.dart';

@riverpod
RunRepository runRepository(RunRepositoryRef ref) {
  return RunRepository(ref.watch(dioProvider));
}

class RunRepository {
  final Dio _dio;

  RunRepository(this._dio);

  Future<RunModel> submitRun({
    required int distance,
    required int duration,
    required String startedAt,
    required String endedAt,
  }) async {
    try {
      final response = await _dio.post(
        ApiConstants.runs,
        data: {
          'distance': distance,
          'duration': duration,
          'started_at': startedAt,
          'ended_at': endedAt,
        },
      );

      if (response.statusCode == 201) {
        return RunModel.fromJson(response.data['data']);
      } else {
        throw Exception('Failed to submit run');
      }
    } on DioException catch (e) {
      throw Exception(e.response?.data['error'] ?? 'Failed to submit run');
    }
  }

  Future<List<RunModel>> getUserRuns() async {
    try {
      final response = await _dio.get(ApiConstants.runs);

      if (response.statusCode == 200) {
        final List<dynamic> data = response.data['data'];
        return data.map((json) => RunModel.fromJson(json)).toList();
      } else {
        throw Exception('Failed to fetch runs');
      }
    } on DioException catch (e) {
      throw Exception(e.response?.data['error'] ?? 'Failed to fetch runs');
    }
  }

  Future<RunStatsModel> getUserStats() async {
    try {
      final response = await _dio.get(ApiConstants.runStats);

      if (response.statusCode == 200) {
        return RunStatsModel.fromJson(response.data['data']);
      } else {
        throw Exception('Failed to fetch stats');
      }
    } on DioException catch (e) {
      throw Exception(e.response?.data['error'] ?? 'Failed to fetch stats');
    }
  }

  Future<RunModel> getRunById(String id) async {
    try {
      final response = await _dio.get('${ApiConstants.runs}/$id');

      if (response.statusCode == 200) {
        return RunModel.fromJson(response.data['data']);
      } else {
        throw Exception('Failed to fetch run');
      }
    } on DioException catch (e) {
      throw Exception(e.response?.data['error'] ?? 'Failed to fetch run');
    }
  }
}

