import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:go_router/go_router.dart';
import 'package:flutter_hooks/flutter_hooks.dart';

import 'package:memoria/ui/components/index.dart';
import 'package:memoria/ui/screens/signup/hooks/use_view_model.dart';

class SignupScreen extends HookWidget {
  final formKey = GlobalKey<FormState>();
  SignupScreen({super.key});

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
                AppLocalizations.of(context)!.w_signup,
                style: const TextStyle(
                  fontSize: 32.0,
                  fontWeight: FontWeight.bold,
                ),
              ),
              const SizedBox(height: 24.0),

              // User name
              TextFormField(
                decoration: InputDecoration(
                  border: const OutlineInputBorder(),
                  labelText: AppLocalizations.of(context)!.w_user_name,
                  suffixIcon: const Icon(
                    Icons.person,
                    size: 24.0,
                  ),
                ),
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return AppLocalizations.of(context)!
                        .v_required('User name');
                  }
                  return null;
                },
                onChanged: (value) {
                  vm.userName.value = value;
                },
              ),
              const SizedBox(height: 16.0),

              // User space name
              TextFormField(
                decoration: InputDecoration(
                  border: const OutlineInputBorder(),
                  labelText: AppLocalizations.of(context)!.w_user_space_name,
                  suffixIcon: const Icon(
                    Icons.home,
                    size: 24.0,
                  ),
                ),
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return AppLocalizations.of(context)!
                        .v_required('User space name');
                  }
                  return null;
                },
                onChanged: (value) {
                  vm.userSpaceName.value = value;
                },
              ),
              const SizedBox(height: 16.0),

              // Email
              TextFormField(
                decoration: InputDecoration(
                  border: const OutlineInputBorder(),
                  labelText: AppLocalizations.of(context)!.w_email,
                  suffixIcon: const Icon(
                    Icons.email,
                    size: 24.0,
                  ),
                ),
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return AppLocalizations.of(context)!.v_required('Email');
                  }
                  return null;
                },
                onChanged: (value) {
                  vm.email.value = value;
                },
              ),
              const SizedBox(height: 16.0),

              // password
              TextFormField(
                decoration: InputDecoration(
                  border: const OutlineInputBorder(),
                  labelText: AppLocalizations.of(context)!.w_password,
                  suffixIcon: const Icon(
                    Icons.password,
                    size: 24.0,
                  ),
                ),
                validator: (value) {
                  if (value == null || value.isEmpty) {
                    return AppLocalizations.of(context)!.v_required('Password');
                  }
                  return null;
                },
                onChanged: (value) {
                  vm.password.value = value;
                },
                obscureText: true,
              ),
              const SizedBox(height: 16.0),

              // Buttons
              SizedBox(
                width: double.infinity,
                child: FilledButton(
                  onPressed: () async {
                    await vm.signup();
                  },
                  child: Text(AppLocalizations.of(context)!.w_signup),
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
