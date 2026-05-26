import 'package:flutter/material.dart';
import 'package:panexpress/ui/home/widget/home_screen.dart';
import 'package:panexpress/ui/pesan_makanan/widget/pesan_makanan_screen.dart';

class AppRoutes {
  static Route<dynamic> generateRoute(
    RouteSettings settings,
  ) {
    switch (settings.name) {
      case '/home':
        return MaterialPageRoute(
          builder: (_) => const HomeScreen(),
        );

      case '/pesan-makanan':
        return MaterialPageRoute(
          builder: (_) => const PesanMakananScreen(),
        );

      default:
        return MaterialPageRoute(
          builder: (_) => Scaffold(
            body: Center(
              child: Text(
                'No route defined for ${settings.name}',
              ),
            ),
          ),
        );
    }
  }
}
