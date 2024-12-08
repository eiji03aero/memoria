class LoginRes {
  final String token;

  LoginRes({required this.token});

  factory LoginRes.fromJson(Map<String, dynamic> json) {
    return LoginRes(token: json['token']);
  }
}
