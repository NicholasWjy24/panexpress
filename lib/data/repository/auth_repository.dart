import 'package:panexpress/data/service/auth_service.dart';

class AuthRepository {
  AuthRepository(this._authService);

  final AuthService _authService;

  Future<String> login({
    required String username,
    required String password,
  }) async {
    final response = await _authService.login(
      username: username,
      password: password,
    );

    return response['message'] ?? 'Login successful';
  }

  Future<String> register({
    required String username,
    required String email,
    required String password,
  }) async {
    final response = await _authService.register(
      username: username,
      email: email,
      password: password,
    );

    return response['message'] ?? 'Registration successful';
  }
}
