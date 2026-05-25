import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:panexpress/data/service/auth_service.dart';
import 'package:panexpress/data/service/menu_service.dart';

import 'data/repository/auth_repository.dart';
import 'data/repository/menu_repository.dart';
import 'ui/login/bloc/auth_bloc.dart';
import 'ui/login/widget/login_screen.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MultiRepositoryProvider(
      providers: [
        RepositoryProvider(
          create: (_) => AuthRepository(AuthService()),
        ),
        RepositoryProvider(
          create: (_) => MenuRepository(MenuService()),
        ),
      ],
      child: BlocProvider(
        create: (context) => AuthBloc(context.read<AuthRepository>()),
        child: MaterialApp(
          title: 'PanExpress',
          debugShowCheckedModeBanner: false,
          theme: ThemeData(
            colorScheme: ColorScheme.fromSeed(seedColor: Colors.teal),
            inputDecorationTheme: const InputDecorationTheme(
              filled: true,
            ),
            useMaterial3: true,
          ),
          home: const LoginScreen(),
        ),
      ),
    );
  }
}
