import 'package:riverpod_annotation/riverpod_annotation.dart';
import '../../../core/models/item_model.dart';
import '../data/shop_repository.dart';

part 'shop_provider.g.dart';

@riverpod
class Shop extends _$Shop {
  @override
  String? build() {
    return null;
  }

  Future<void> buyItem(String itemId) async {
    final shopRepository = ref.read(shopRepositoryProvider);
    await shopRepository.buyItem(itemId);
    state = itemId;
  }

  Future<void> equipItem(String itemId) async {
    final shopRepository = ref.read(shopRepositoryProvider);
    await shopRepository.equipItem(itemId);
  }

  Future<void> unequipItem(String itemId) async {
    final shopRepository = ref.read(shopRepositoryProvider);
    await shopRepository.unequipItem(itemId);
  }
}

@riverpod
Future<List<ItemModel>> shopItems(ShopItemsRef ref) async {
  final shopRepository = ref.watch(shopRepositoryProvider);
  return await shopRepository.getAvailableItems();
}

