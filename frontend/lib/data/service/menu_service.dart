import 'dart:convert';

import 'package:http/http.dart' as http;

class MenuService {
  MenuService({
    http.Client? client,
    this.baseUrl = 'http://10.0.2.2:8080/api',
  }) : _client = client ?? http.Client();

  final http.Client _client;
  final String baseUrl;

  Future<dynamic> getMenus() async {
    final response = await _client.get(
      Uri.parse('$baseUrl/menu'),
      headers: const {
        'Accept': 'application/json',
      },
    );

    final decoded = response.body.isNotEmpty ? jsonDecode(response.body) : [];

    if (response.statusCode >= 200 && response.statusCode < 300) {
      return decoded;
    }

    if (decoded is Map<String, dynamic>) {
      throw MenuException(
        decoded['error'] ?? 'Request failed with status ${response.statusCode}',
      );
    }

    throw MenuException('Request failed with status ${response.statusCode}');
  }
}

class MenuException implements Exception {
  const MenuException(this.message);

  final String message;

  @override
  String toString() => message;
}
