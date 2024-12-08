import 'package:memoria/data/repositories/account_repository.dart';
import 'package:memoria/data/services/secret_service.dart';

class AccountUsecase {
  final AccountRepository _accountRepo;

  AccountUsecase() : _accountRepo = AccountRepository();

  Future<void> signup({
    required String userName,
    required String userSpaceName,
    required String email,
    required String password,
  }) async {
    await _accountRepo.signup(
      userName: userName,
      userSpaceName: userSpaceName,
      email: email,
      password: password,
    );
  }

  Future<void> login({
    required String email,
    required String password,
  }) async {
    final token = await _accountRepo.login(
      email: email,
      password: password,
    );

    final secretSvc = await SecretService.init();
    await secretSvc.saveApiToken(token);
  }
}
