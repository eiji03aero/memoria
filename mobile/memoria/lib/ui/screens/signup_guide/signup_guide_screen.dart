import 'package:flutter/material.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:go_router/go_router.dart';
import 'package:flutter_hooks/flutter_hooks.dart';

import 'package:memoria/ui/components/index.dart';

class SignupGuideScreen extends HookWidget {
  const SignupGuideScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: centeredBox(
        maxWidth: 340.0,
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            Text(
              AppLocalizations.of(context)!.s_welcome_guide_h_1,
              style: const TextStyle(
                fontSize: 20.0,
                fontWeight: FontWeight.bold,
              ),
            ),
            const SizedBox(height: 16.0),

            Text(
              AppLocalizations.of(context)!.s_welcome_guide_p_1,
              style: const TextStyle(
                fontSize: 16.0,
              ),
            ),
            const SizedBox(height: 16.0),

            // Buttons
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
    );
  }
}
