import 'package:dio/dio.dart';
import 'dart:developer' as developer;

import 'package:memoria/config.dart';

class HttpClient {
  final Dio client;
  HttpClient({required this.client});

  factory HttpClient.init() {
    return HttpClient(client: Dio());
  }

  String getMeoriaApiUrl(String path) {
    return "${Config.memoriaApiHost()}$path";
  }

  Future<Response<T>> post<T>({required String url, Object? data}) async {
    developer.log("Starting HttpClient#post $url $data", name: "HttpClient");
    final res = await client.post<T>(url, data: data);
    developer.log("Finished HttpClient#post $url $data", name: "HttpClient");
    return res;
  }
}
