import 'package:flutter/material.dart';
import 'package:flutter_localizations/flutter_localizations.dart';
import 'package:flutter_gen/gen_l10n/app_localizations.dart';
import 'package:hive_flutter/hive_flutter.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';

import 'package:memoria/ui/router.dart';
import 'package:memoria/config.dart';

void main() async {
  await Hive.initFlutter();
  await Config.init();
  runApp(const ProviderScope(child: Memoria()));
}

class Memoria extends StatelessWidget {
  const Memoria({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
      title: 'Memoria',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(
          seedColor: const Color(0xFFFC3E1B),
          primary: const Color(0xFFFC3E1B),
        ),
        useMaterial3: true,
      ),
      localizationsDelegates: const [
        AppLocalizations.delegate,
        GlobalMaterialLocalizations.delegate,
        GlobalWidgetsLocalizations.delegate,
        GlobalCupertinoLocalizations.delegate,
      ],
      supportedLocales: const [
        Locale('en'),
        Locale('ja'),
      ],
      routerConfig: goRouter,
    );
  }
}
