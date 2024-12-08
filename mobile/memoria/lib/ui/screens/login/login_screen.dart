import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:go_router/go_router.dart';
import 'package:flutter_hooks/flutter_hooks.dart';

import 'package:memoria/ui/components/index.dart';
import 'package:memoria/ui/screens/login/hooks/use_view_model.dart';

class LoginScreen extends HookWidget {
  final formKey = GlobalKey<FormState>();
  LoginScreen({super.key});

  @override
  Widget build(BuildContext context) {
    final vm = useViewModel();

    return Scaffold(
      body: centeredBox(
        maxWidth: 340.0,
        child: Form(
          key: formKey,
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: <Widget>[
              Text(
                AppLocalizations.of(context)!.w_login,
                style: const TextStyle(
                  fontSize: 32.0,
                  fontWeight: FontWeight.bold,
                ),
              ),
              const SizedBox(height: 20.0),
              TextFormField(
                  decoration: InputDecoration(
                    border: const OutlineInputBorder(),
                    labelText: AppLocalizations.of(context)!.w_email,
                    suffixIcon: const Icon(
                      Icons.email,
                      size: 16.0,
                    ),
                  ),
                  onChanged: (value) {
                    vm.email.value = value;
                  }),
              const SizedBox(height: 20.0),
              TextFormField(
                decoration: InputDecoration(
                  border: const OutlineInputBorder(),
                  labelText: AppLocalizations.of(context)!.w_password,
                  suffixIcon: const Icon(
                    Icons.password,
                    size: 16.0,
                  ),
                ),
                onChanged: (value) {
                  vm.password.value = value;
                },
                obscureText: true,
              ),
              const SizedBox(height: 20.0),
              SizedBox(
                width: double.infinity,
                child: FilledButton(
                  onPressed: vm.login,
                  child: Text(AppLocalizations.of(context)!.w_login),
                ),
              ),
              const SizedBox(height: 24.0),
              dividerLabel(child: Text(AppLocalizations.of(context)!.w_or)),
              const SizedBox(height: 24.0),
              SizedBox(
                width: double.infinity,
                child: TextButton(
                  onPressed: () {
                    context.go("/welcome");
                  },
                  child: Text(AppLocalizations.of(context)!.w_go_back),
                ),
              )
            ],
          ),
        ),
      ),
    );
  }
}
