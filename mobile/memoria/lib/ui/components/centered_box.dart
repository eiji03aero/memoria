import 'package:flutter/material.dart';

Widget centeredBox({required Widget child, double maxWidth = 300.0}) {
  return Container(
    padding: const EdgeInsets.all(16.0),
    child: Center(
      child: Align(
        alignment: Alignment.center,
        child: ConstrainedBox(
          constraints: BoxConstraints(maxWidth: maxWidth),
          child: child,
        ),
      ),
    ),
  );
}
