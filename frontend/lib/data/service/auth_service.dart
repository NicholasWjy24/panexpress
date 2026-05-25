import 'dart:async';
import 'dart:convert';
import 'dart:developer';
import 'dart:io';

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
    try {
      log(
        'REQUEST → POST $url',
        name: 'AuthService',
      );

      final response = await _client
          .post(
            Uri.parse(url),
            headers: const {
              'Accept': 'application/json',
              'Content-Type': 'application/json',
            },
            body: jsonEncode(body),
          )
          .timeout(const Duration(seconds: 10));

      log(
        'RESPONSE CODE : ${response.statusCode}\nRESPONSE MESSAGE: ${response.body}',
        name: 'AuthService',
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
    } on SocketException catch (e, stackTrace) {
      log(
        'SocketException: Cannot connect to server.',
        name: 'AuthService',
        error: e,
        stackTrace: stackTrace,
      );

      throw const AuthException(
        'Cannot connect to server.',
      );
    } on TimeoutException catch (e, stackTrace) {
      log(
        'TimeoutException: Connection timeout.',
        name: 'AuthService',
        error: e,
        stackTrace: stackTrace,
      );

      throw const AuthException(
        'Connection timeout.',
      );
    } on FormatException catch (e, stackTrace) {
      log(
        'FormatException: Invalid server response.',
        name: 'AuthService',
        error: e,
        stackTrace: stackTrace,
      );

      throw const AuthException(
        'Invalid server response.',
      );
    } on http.ClientException catch (e, stackTrace) {
      log(
        'ClientException: HTTP client error.',
        name: 'AuthService',
        error: e,
        stackTrace: stackTrace,
      );

      throw const AuthException(
        'HTTP client error.',
      );
    } catch (e, stackTrace) {
      log(
        'Unexpected Error',
        name: 'AuthService',
        error: e,
        stackTrace: stackTrace,
      );

      throw AuthException(
        'Unexpected error: $e',
      );
    }
  }
}

class AuthException implements Exception {
  const AuthException(this.message);

  final String message;

  @override
  String toString() => message;
}
