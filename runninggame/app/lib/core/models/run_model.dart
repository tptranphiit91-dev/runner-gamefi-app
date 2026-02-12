import 'package:freezed_annotation/freezed_annotation.dart';

part 'run_model.freezed.dart';
part 'run_model.g.dart';

@freezed
class RunModel with _$RunModel {
  const factory RunModel({
    required String id,
    @JsonKey(name: 'user_id') required String userId,
    required int distance,
    required int duration,
    @JsonKey(name: 'started_at') required String startedAt,
    @JsonKey(name: 'ended_at') required String endedAt,
    @JsonKey(name: 'earned_coins') required double earnedCoins,
    @JsonKey(name: 'created_at') required String createdAt,
  }) = _RunModel;

  factory RunModel.fromJson(Map<String, dynamic> json) =>
      _$RunModelFromJson(json);
}

@freezed
class RunStatsModel with _$RunStatsModel {
  const factory RunStatsModel({
    @JsonKey(name: 'total_distance') required int totalDistance,
    @JsonKey(name: 'total_duration') required int totalDuration,
    @JsonKey(name: 'total_earned_coins') required double totalEarnedCoins,
    @JsonKey(name: 'total_runs') required int totalRuns,
    @JsonKey(name: 'avg_distance') required double avgDistance,
    @JsonKey(name: 'avg_duration') required double avgDuration,
  }) = _RunStatsModel;

  factory RunStatsModel.fromJson(Map<String, dynamic> json) =>
      _$RunStatsModelFromJson(json);
}

