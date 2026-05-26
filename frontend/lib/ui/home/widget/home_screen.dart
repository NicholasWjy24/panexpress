import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:panexpress/data/model/menu_item.dart';
import 'package:panexpress/data/repository/menu_repository.dart';
import 'package:panexpress/ui/home/bloc/home_bloc.dart';
import 'package:panexpress/ui/login/bloc/auth_bloc.dart';
import 'package:panexpress/ui/login/widget/login_screen.dart';

class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (context) => HomeBloc(
        context.read<MenuRepository>(),
      )..add(const HomeStarted()),
      child: const _HomeView(),
    );
  }
}

class _HomeView extends StatelessWidget {
  const _HomeView();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Home'),
      ),
      drawer: const _HomeDrawer(),
      body: const Center(
        child: Text('Home Screen'),
      ),
    );
  }
}

class _HomeDrawer extends StatelessWidget {
  const _HomeDrawer();

  @override
  Widget build(BuildContext context) {
    return Drawer(
      child: SafeArea(
        child: Column(
          children: [
            DrawerHeader(
              margin: EdgeInsets.zero,
              child: Align(
                alignment: Alignment.centerLeft,
                child: Text(
                  'PanExpress',
                  style: Theme.of(context).textTheme.headlineSmall,
                ),
              ),
            ),
            Expanded(
              child: BlocBuilder<HomeBloc, HomeState>(
                builder: (context, state) {
                  if (state is HomeLoading || state is HomeInitial) {
                    return const Center(
                      child: CircularProgressIndicator(),
                    );
                  }

                  if (state is HomeFailure) {
                    return _DrawerMessage(
                      icon: Icons.error_outline,
                      message: state.message,
                      actionLabel: 'Retry',
                      onActionPressed: () {
                        context
                            .read<HomeBloc>()
                            .add(const HomeMenuRefreshRequested());
                      },
                    );
                  }

                  if (state is HomeSuccess && state.menus.isEmpty) {
                    return const _DrawerMessage(
                      icon: Icons.menu_open_outlined,
                      message: 'No menus available.',
                    );
                  }

                  if (state is HomeSuccess) {
                    return ListView.builder(
                      padding: EdgeInsets.zero,
                      itemCount: state.menus.length,
                      itemBuilder: (context, index) {
                        final menu = state.menus[index];

                        return ListTile(
                          leading: const Icon(Icons.chevron_right),
                          title: Text(menu.menuName),
                          onTap: () => _selectMenu(context, menu),
                        );
                      },
                    );
                  }

                  return const SizedBox.shrink();
                },
              ),
            ),
            const Divider(height: 1),
            ListTile(
              leading: const Icon(Icons.logout_outlined),
              title: const Text('Logout'),
              onTap: () => _logout(context),
            ),
          ],
        ),
      ),
    );
  }

  void _selectMenu(BuildContext context, MenuItem menu) {
    Navigator.pop(context);

    if (menu.route.isNotEmpty) {
      Navigator.pushNamed(
        context,
        menu.route,
      );
    }
  }

  void _logout(BuildContext context) {
    context.read<AuthBloc>().add(const LogoutRequested());

    Navigator.of(context).pushAndRemoveUntil(
      MaterialPageRoute(
        builder: (_) => const LoginScreen(),
      ),
      (_) => false,
    );
  }
}

class _DrawerMessage extends StatelessWidget {
  const _DrawerMessage({
    required this.icon,
    required this.message,
    this.actionLabel,
    this.onActionPressed,
  });

  final IconData icon;
  final String message;
  final String? actionLabel;
  final VoidCallback? onActionPressed;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(24),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Icon(
            icon,
            size: 40,
            color: Theme.of(context).colorScheme.outline,
          ),
          const SizedBox(height: 12),
          Text(
            message,
            textAlign: TextAlign.center,
          ),
          if (actionLabel != null && onActionPressed != null) ...[
            const SizedBox(height: 12),
            TextButton(
              onPressed: onActionPressed,
              child: Text(actionLabel!),
            ),
          ],
        ],
      ),
    );
  }
}
