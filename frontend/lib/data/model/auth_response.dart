class AuthResponse {
  const AuthResponse({
    required this.message,
    required this.token,
    required this.userId,
    required this.username,
    required this.email,
    required this.roleLevel,
  });

  final String message;
  final String token;
  final int userId;
  final String username;
  final String email;
  final int roleLevel;

  factory AuthResponse.fromJson(Map<String, dynamic> json) {
    return AuthResponse(
      message: json['message'] ?? '',
      token: json['token'] ?? '',
      userId: json['user_id'] ?? 0,
      username: json['username'] ?? '',
      email: json['email'] ?? '',
      roleLevel: json['role_level'] ?? 3,
    );
  }
}
