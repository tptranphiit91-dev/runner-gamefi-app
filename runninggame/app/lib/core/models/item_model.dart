import 'package:freezed_annotation/freezed_annotation.dart';

part 'item_model.freezed.dart';
part 'item_model.g.dart';

@freezed
class ItemModel with _$ItemModel {
  const factory ItemModel({
    required String id,
    required String name,
    required String type,
    required double price,
    @JsonKey(name: 'stat_multiplier') required double statMultiplier,
    required String description,
    @JsonKey(name: 'image_url') required String imageUrl,
    @JsonKey(name: 'is_available') required bool isAvailable,
  }) = _ItemModel;

  factory ItemModel.fromJson(Map<String, dynamic> json) =>
      _$ItemModelFromJson(json);
}

