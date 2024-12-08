import 'package:flutter_dotenv/flutter_dotenv.dart';

class Config {
  static Future<void> init() async {
    await dotenv.load(fileName: ".env");
  }

  static String _getEnvVar(String key) {
    final value = dotenv.env[key];
    if (value == null || value.isEmpty) {
      throw Exception("env variable not set: $key");
    }

    return value;
  }

  static String memoriaApiHost() {
    return _getEnvVar("MEMORIA_API_HOST");
  }
}
