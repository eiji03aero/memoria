import 'package:memoria/data/models/login.dart';
import 'package:memoria/data/services/http_client.dart';

class AccountService {
  final HttpClient _httpClient;
  AccountService() : _httpClient = HttpClient.init();

  Future<void> signup({
    required String userName,
    required String userSpaceName,
    required String email,
    required String password,
  }) async {
    await _httpClient
        .post(url: _httpClient.getMeoriaApiUrl('/api/public/signup'), data: {
      'name': userName,
      'user_space_name': userSpaceName,
      'email': email,
      'password': password,
    });
  }

  Future<LoginRes> login({
    required String email,
    required String password,
  }) async {
    final res = await _httpClient
        .post(url: _httpClient.getMeoriaApiUrl('/api/public/login'), data: {
      'email': email,
      'password': password,
    });
    return LoginRes.fromJson(res.data);
  }
}
