import 'package:bloc_test/bloc_test.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mocktail/mocktail.dart';
import 'package:panexpress/data/model/auth_response.dart';
import 'package:panexpress/data/repository/auth_repository.dart';
import 'package:panexpress/data/service/auth_service.dart';
import 'package:panexpress/ui/login/bloc/auth_bloc.dart';

class MockAuthRepository extends Mock implements AuthRepository {}

void main() {
  late MockAuthRepository mockAuthRepository;
  late AuthBloc authBloc;

  setUp(() {
    mockAuthRepository = MockAuthRepository();
    authBloc = AuthBloc(mockAuthRepository);
  });

  tearDown(() {
    authBloc.close();
  });

  test('initial state harus berupa AuthInitial', () {
    print('\n[RUNNING TEST] → Mengecek Initial State');
    expect(authBloc.state, const AuthInitial());
    print('✅ [SUCCESS] → Initial State sesuai (AuthInitial)');
  });

  group('LoginSubmitted Event Tests', () {
    const tUsername = 'admin';
    const tPassword = 'Addminn7';

    const ntUsername = 'aaa';
    const ntPassword = 'aaa';

    // TEST 2: Skenario Login Berhasil
    blocTest<AuthBloc, AuthState>(
      'harus memancarkan [AuthLoading, AuthSuccess] ketika login di repository berhasil',
      setUp: () {
        print('\n[RUNNING TEST] → Skenario: Login Berhasil');
      },
      build: () {
        when(
          () => mockAuthRepository.login(
            username: tUsername,
            password: tPassword,
          ),
        ).thenAnswer(
          (_) async => const AuthResponse(
            message: 'Login Berhasil!',
            token: 'dummy-token',
            userId: 1,
            username: 'nick',
            email: 'nick@gmail.com',
            roleLevel: 1,
          ),
        );

        return authBloc;
      },
      act: (bloc) => bloc.add(
        const LoginSubmitted(
          username: tUsername,
          password: tPassword,
        ),
      ),
      expect: () => [
        const AuthLoading(),
        const LoginAuthSuccess(
          AuthResponse(
            message: 'Login Berhasil!',
            token: 'dummy-token',
            userId: 1,
            username: 'nick',
            email: 'nick@gmail.com',
            roleLevel: 1,
          ),
        ),
      ],
      verify: (_) {
        verify(
          () => mockAuthRepository.login(
            username: tUsername,
            password: tPassword,
          ),
        ).called(1);

        print(
          '✅ [SUCCESS] → Skenario Login Berhasil Lolos! State sesuai ekspektasi.',
        );
      },
    );

    // TEST 3: Skenario Login Gagal karena salah password/username (AuthException)
    // Contoh modifikasi TEST 3 di file test kamu tanpa Equatable:
    blocTest<AuthBloc, AuthState>(
      'harus memancarkan [AuthLoading, AuthFailure] ketika terjadi AuthException',
      build: () {
        when(() => mockAuthRepository.login(
                username: ntUsername, password: ntPassword))
            .thenThrow(const AuthException('Username atau password salah!'));
        return authBloc;
      },
      act: (bloc) => bloc.add(
          const LoginSubmitted(username: ntUsername, password: ntPassword)),
      expect: () => [
        const AuthLoading(),
        // Menggunakan Matchers "isA" untuk mengecek tipenya, lalu mengecek isinya
        isA<AuthFailure>().having((state) => state.message, 'message',
            'Username atau password salah!'),
      ],
      verify: (_) {
        print(
            '✅ [SUCCESS] → Skenario Username / Password salah berhasil! Pesan fallback berhasil muncul.');
      },
    );

    // TEST 4: Skenario Server Backend Mati atau Internet Terputus (Generic Exception)
    blocTest<AuthBloc, AuthState>(
      'harus memancarkan [AuthLoading, AuthFailure] dengan pesan fallback ketika server tidak merespon',
      setUp: () {
        print(
            '\n[RUNNING TEST] → Skenario: Server Down/Timeout (Generic Exception)');
      },
      build: () {
        when(() => mockAuthRepository.login(
              username: tUsername,
              password: tPassword,
            )).thenThrow(Exception('Server Timeout'));
        return authBloc;
      },
      act: (bloc) => bloc
          .add(const LoginSubmitted(username: tUsername, password: tPassword)),
      expect: () => [
        const AuthLoading(),
        const AuthFailure('Could not connect to the login API.'),
      ],
      verify: (_) {
        print(
            '✅ [SUCCESS] → Skenario Generic Exception Lolos! Pesan fallback berhasil muncul.');
      },
    );
  });
}
