import 'package:panexpress/data/model/menu_item.dart';
import 'package:panexpress/data/service/menu_service.dart';

class MenuRepository {
  MenuRepository(this._menuService);

  final MenuService _menuService;

  Future<List<MenuItem>> getMenus() async {
    final response = await _menuService.getMenus();
    final rawMenus = _extractMenuList(response);

    return rawMenus
        .whereType<Map<String, dynamic>>()
        .map(MenuItem.fromJson)
        .toList();
  }

  List<dynamic> _extractMenuList(dynamic response) {
    if (response is List) {
      return response;
    }

    if (response is Map<String, dynamic>) {
      final candidates = [
        response['menus'],
        response['menu'],
        response['data'],
        response['items'],
      ];

      for (final candidate in candidates) {
        if (candidate is List) {
          return candidate;
        }
      }
    }

    return const [];
  }
}
