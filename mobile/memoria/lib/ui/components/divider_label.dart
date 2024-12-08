import 'package:flutter/material.dart';

Widget dividerLabel({required Widget child}) {
  return Row(children: [
    Expanded(
      child: Container(height: 1.0, color: Colors.grey[200]),
    ),
    const SizedBox(width: 24.0),
    child,
    const SizedBox(width: 24.0),
    Expanded(
      child: Container(height: 1.0, color: Colors.grey[200]),
    ),
  ]);
}
