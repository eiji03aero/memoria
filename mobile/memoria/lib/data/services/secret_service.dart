import 'package:hive_flutter/hive_flutter.dart';

class SecretService {
  late final Box<dynamic> box;
  SecretService._internal({required this.box});

  static Future<SecretService> init() async {
    final box = await Hive.openBox('secret');
    return SecretService._internal(box: box);
  }

  Future<void> saveApiToken(String token) async {
    await box.put('apiToken', token);
  }

  Future<String?> getApiToken() async {
    return box.get('apiToken');
  }
}
