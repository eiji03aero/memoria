import 'package:memoria/data/services/account_service.dart';

class AccountRepository {
  final AccountService _accountService;
  AccountRepository() : _accountService = AccountService();

  Future<void> signup({
    required String userName,
    required String userSpaceName,
    required String email,
    required String password,
  }) async {
    await _accountService.signup(
      userName: userName,
      userSpaceName: userSpaceName,
      email: email,
      password: password,
    );
  }

  Future<String> login({
    required String email,
    required String password,
  }) async {
    final res = await _accountService.login(
      email: email,
      password: password,
    );
    return res.token;
  }
}
