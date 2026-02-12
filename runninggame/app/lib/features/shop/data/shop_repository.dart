import 'package:dio/dio.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';
import '../../../core/constants/api_constants.dart';
import '../../../core/models/item_model.dart';
import '../../../core/network/dio_client.dart';

part 'shop_repository.g.dart';

@riverpod
ShopRepository shopRepository(ShopRepositoryRef ref) {
  return ShopRepository(ref.watch(dioProvider));
}

class ShopRepository {
  final Dio _dio;

  ShopRepository(this._dio);

  Future<List<ItemModel>> getAvailableItems() async {
    try {
      final response = await _dio.get(ApiConstants.items);

      if (response.statusCode == 200) {
        final List<dynamic> data = response.data['data'];
        return data.map((json) => ItemModel.fromJson(json)).toList();
      } else {
        throw Exception('Failed to fetch items');
      }
    } on DioException catch (e) {
      throw Exception(e.response?.data['error'] ?? 'Failed to fetch items');
    }
  }

  Future<void> buyItem(String itemId) async {
    try {
      final response = await _dio.post(
        ApiConstants.buyItem,
        data: {'item_id': itemId},
      );

      if (response.statusCode != 200) {
        throw Exception('Failed to purchase item');
      }
    } on DioException catch (e) {
      throw Exception(e.response?.data['error'] ?? 'Failed to purchase item');
    }
  }

  Future<void> equipItem(String itemId) async {
    try {
      final response = await _dio.post(
        ApiConstants.equipItem,
        data: {'item_id': itemId},
      );

      if (response.statusCode != 200) {
        throw Exception('Failed to equip item');
      }
    } on DioException catch (e) {
      throw Exception(e.response?.data['error'] ?? 'Failed to equip item');
    }
  }

  Future<void> unequipItem(String itemId) async {
    try {
      final response = await _dio.post(
        ApiConstants.unequipItem,
        data: {'item_id': itemId},
      );

      if (response.statusCode != 200) {
        throw Exception('Failed to unequip item');
      }
    } on DioException catch (e) {
      throw Exception(e.response?.data['error'] ?? 'Failed to unequip item');
    }
  }
}

