import 'package:flutter_test/flutter_test.dart';

import 'package:panexpress/main.dart';

void main() {
  testWidgets('shows login form and can switch to register', (tester) async {
    await tester.pumpWidget(const MyApp());

    expect(find.text('Welcome back'), findsOneWidget);
    expect(find.text('Login'), findsOneWidget);
    expect(find.text('Email'), findsOneWidget);
    expect(find.text('Password'), findsOneWidget);

    await tester.tap(find.text('Need an account? Register'));
    await tester.pump();

    expect(find.text('Create account'), findsOneWidget);
    expect(find.text('Name'), findsOneWidget);
    expect(find.text('Register'), findsOneWidget);
  });
}
