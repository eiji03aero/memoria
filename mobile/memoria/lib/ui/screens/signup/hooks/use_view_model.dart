import 'package:flutter/foundation.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:go_router/go_router.dart';

import 'package:memoria/domain/usecases/account_usecase.dart';

ViewModel useViewModel() {
  final context = useContext();
  final isMutating = useState(false);

  final userName = useState('');
  final userSpaceName = useState('');
  final email = useState('');
  final password = useState('');

  final accountUc = AccountUsecase();

  Future<void> signup() async {
    isMutating.value = true;

    await accountUc.signup(
      userName: userName.value,
      userSpaceName: userSpaceName.value,
      email: email.value,
      password: password.value,
    );

    isMutating.value = false;

    if (context.mounted) {
      context.go('/signup-guide');
    }
  }

  return ViewModel(
    signup: signup,
    isMutating: isMutating,
    userName: userName,
    userSpaceName: userSpaceName,
    email: email,
    password: password,
  );
}

class ViewModel {
  final Future<void> Function() signup;
  final ValueNotifier<bool> isMutating;
  final ValueNotifier<String> userName;
  final ValueNotifier<String> userSpaceName;
  final ValueNotifier<String> email;
  final ValueNotifier<String> password;

  ViewModel({
    required this.signup,
    required this.isMutating,
    required this.userName,
    required this.userSpaceName,
    required this.email,
    required this.password,
  });
}
