import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:go_router/go_router.dart';

import 'package:memoria/ui/components/index.dart';

class WelcomeScreen extends StatelessWidget {
  const WelcomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: centeredBox(
        maxWidth: 340.0,
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            Image.asset(
              'assets/welcome/bg.png',
              height: 200.0,
            ),
            const SizedBox(height: 20.0),
            Text(
              AppLocalizations.of(context)!.s_welcome_heading_1,
              style: const TextStyle(
                fontSize: 24.0,
                fontWeight: FontWeight.bold,
              ),
            ),
            Text(
              AppLocalizations.of(context)!.s_welcome_heading_2,
              style: const TextStyle(
                fontSize: 24.0,
                fontWeight: FontWeight.bold,
              ),
            ),
            const SizedBox(height: 20.0),
            SizedBox(
              width: double.infinity,
              child: FilledButton(
                onPressed: () {
                  context.push('/login');
                },
                child: Text(AppLocalizations.of(context)!.w_login),
              ),
            ),
            const SizedBox(height: 20.0),
            SizedBox(
              width: double.infinity,
              child: TextButton(
                onPressed: () {
                  context.push('/signup');
                },
                child: Text(AppLocalizations.of(context)!.w_signup),
              ),
            )
          ],
        ),
      ),
    );
  }
}
