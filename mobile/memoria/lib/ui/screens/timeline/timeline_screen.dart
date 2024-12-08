import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';

import 'package:memoria/ui/components/index.dart';

class TimelineScreen extends HookWidget {
  const TimelineScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: centeredBox(
        maxWidth: 340.0,
        child: const Text(
          "I'm gonna go, go. Are you gonna go?",
          style: TextStyle(
            fontSize: 32.0,
            fontWeight: FontWeight.bold,
          ),
        ),
      ),
    );
  }
}
