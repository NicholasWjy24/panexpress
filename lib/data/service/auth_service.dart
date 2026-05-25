import 'dart:convert';

import 'package:http/http.dart' as http;

class AuthService {
  AuthService({
    http.Client? client,
    this.baseUrl = 'http://10.0.2.2:8080/api',
  }) : _client = client ?? http.Client();

  final http.Client _client;
  final String baseUrl;

  Future<Map<String, dynamic>> login({
    required String username,
    required String password,
  }) async {
    return _post(
      '$baseUrl/login',
      body: {
        'username': username,
        'password': password,
      },
    );
  }

  Future<Map<String, dynamic>> register({
    required String username,
    required String email,
    required String password,
  }) async {
    return _post(
      '$baseUrl/register',
      body: {
        'username': username,
        'email': email,
        'password': password,
      },
    );
  }

  Future<Map<String, dynamic>> _post(
    String url, {
    required Map<String, String> body,
  }) async {
    final response = await _client.post(
      Uri.parse(url),
      headers: const {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
      },
      body: jsonEncode(body),
    );

    final decoded = response.body.isNotEmpty
        ? jsonDecode(response.body)
        : <String, dynamic>{};

    if (response.statusCode >= 200 && response.statusCode < 300) {
      return decoded;
    }

    throw AuthException(
      decoded['error'] ?? 'Request failed with status ${response.statusCode}',
    );
  }
}

class AuthException implements Exception {
  const AuthException(this.message);

  final String message;

  @override
  String toString() => message;
}
