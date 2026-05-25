import 'package:equatable/equatable.dart';

class MenuItem extends Equatable {
  const MenuItem({
    required this.id,
    required this.menuName,
    required this.minRoleLevel,
    required this.maxRoleLevel,
  });

  final int id;
  final String menuName;
  final int minRoleLevel;
  final int maxRoleLevel;

  factory MenuItem.fromJson(Map<String, dynamic> json) {
    return MenuItem(
      id: json['id'] as int,
      menuName: json['menu_name'] as String,
      minRoleLevel: json['min_role_level'] as int,
      maxRoleLevel: json['max_role_level'] as int,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'menu_name': menuName,
      'min_role_level': minRoleLevel,
      'max_role_level': maxRoleLevel,
    };
  }

  @override
  List<Object?> get props => [
        id,
        menuName,
        minRoleLevel,
        maxRoleLevel,
      ];
}
