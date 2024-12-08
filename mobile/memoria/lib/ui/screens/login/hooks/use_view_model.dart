import 'package:flutter/foundation.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:go_router/go_router.dart';

import 'package:memoria/domain/usecases/account_usecase.dart';

ViewModel useViewModel() {
  final context = useContext();
  final isMutating = useState(false);

  final email = useState('');
  final password = useState('');

  final accountUc = AccountUsecase();

  Future<void> login() async {
    isMutating.value = true;

    await accountUc.login(email: email.value, password: password.value);

    isMutating.value = false;

    if (context.mounted) {
      context.go('/timeline');
    }
  }

  return ViewModel(
    isMutating: isMutating,
    email: email,
    password: password,
    login: login,
  );
}

class ViewModel {
  final ValueNotifier<bool> isMutating;
  final ValueNotifier<String> email;
  final ValueNotifier<String> password;
  final Future<void> Function() login;

  ViewModel({
    required this.isMutating,
    required this.email,
    required this.password,
    required this.login,
  });
}
