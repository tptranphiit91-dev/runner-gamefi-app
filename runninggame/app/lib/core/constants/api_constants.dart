class ApiConstants {
  // Base URL - Update this to your backend URL
  static const String baseUrl = 'http://192.168.1.33:8080';
  static const String apiVersion = '/api/v1';
  
  // Auth Endpoints
  static const String register = '$apiVersion/auth/register';
  static const String login = '$apiVersion/auth/login';
  
  // Run Endpoints
  static const String runs = '$apiVersion/runs';
  static const String runStats = '$apiVersion/runs/stats';
  
  // Item/Shop Endpoints
  static const String items = '$apiVersion/items';
  static const String buyItem = '$apiVersion/items/buy';
  static const String equipItem = '$apiVersion/items/equip';
  static const String unequipItem = '$apiVersion/items/unequip';
  
  // Headers
  static const String authHeader = 'Authorization';
  static const String bearerPrefix = 'Bearer';
  
  // Timeouts
  static const Duration connectTimeout = Duration(seconds: 30);
  static const Duration receiveTimeout = Duration(seconds: 30);
}

