import 'dart:io';
import 'package:dio/dio.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';
import '../../../core/constants/api_constants.dart';

part 'avatar_service.g.dart';

/// AI Avatar Generation Service
/// Supports multiple AI providers:
/// 1. Custom Backend API (Python FastAPI + Stable Diffusion)
/// 2. Replicate API (img2img models)
/// 3. OpenAI DALL-E API
/// 4. Stability AI API
@riverpod
AvatarService avatarService(AvatarServiceRef ref) {
  return AvatarService();
}

class AvatarService {
  final Dio _dio = Dio();

  /// Generate 2D avatar from photo using AI
  /// 
  /// Options:
  /// - style: 'cartoon', 'anime', 'pixel', 'realistic'
  /// - provider: 'custom', 'replicate', 'openai', 'stability'
  Future<AvatarResult> generateAvatar(
    File imageFile, {
    String style = 'cartoon',
    String provider = 'custom',
  }) async {
    switch (provider) {
      case 'replicate':
        return _generateWithReplicate(imageFile, style);
      case 'openai':
        return _generateWithOpenAI(imageFile, style);
      case 'stability':
        return _generateWithStability(imageFile, style);
      case 'custom':
      default:
        return _generateWithCustomBackend(imageFile, style);
    }
  }

  /// Option 1: Custom Backend (Recommended)
  /// Uses Stable Diffusion or similar models
  Future<AvatarResult> _generateWithCustomBackend(
    File imageFile,
    String style,
  ) async {
    try {
      final formData = FormData.fromMap({
        'image': await MultipartFile.fromFile(
          imageFile.path,
          filename: 'avatar.jpg',
        ),
        'style': style,
      });

      final response = await _dio.post(
        '${ApiConstants.baseUrl}/api/v1/avatar/generate',
        data: formData,
        options: Options(
          headers: {
            'Content-Type': 'multipart/form-data',
          },
        ),
      );

      return AvatarResult(
        imageUrl: response.data['avatar_url'],
        thumbnailUrl: response.data['thumbnail_url'],
        style: style,
      );
    } catch (e) {
      throw AvatarGenerationException('Failed to generate avatar: $e');
    }
  }

  /// Option 2: Replicate API
  /// Uses models like:
  /// - "jagilley/controlnet-canny" (edge detection + style transfer)
  /// - "tencentarc/gfpgan" (face restoration)
  /// - "cjwbw/anything-v3.0" (anime style)
  Future<AvatarResult> _generateWithReplicate(
    File imageFile,
    String style,
  ) async {
    const apiToken = String.fromEnvironment('REPLICATE_API_TOKEN');
    
    try {
      // Step 1: Upload image
      final imageBytes = await imageFile.readAsBytes();
      final base64Image = 'data:image/jpeg;base64,${imageBytes.toString()}';

      // Step 2: Create prediction
      final response = await _dio.post(
        'https://api.replicate.com/v1/predictions',
        data: {
          'version': _getReplicateModelVersion(style),
          'input': {
            'image': base64Image,
            'prompt': _getStylePrompt(style),
            'num_outputs': 1,
          },
        },
        options: Options(
          headers: {
            'Authorization': 'Token $apiToken',
            'Content-Type': 'application/json',
          },
        ),
      );

      final predictionId = response.data['id'];

      // Step 3: Poll for result
      return await _pollReplicateResult(predictionId, apiToken);
    } catch (e) {
      throw AvatarGenerationException('Replicate API error: $e');
    }
  }

  /// Option 3: OpenAI DALL-E API
  Future<AvatarResult> _generateWithOpenAI(
    File imageFile,
    String style,
  ) async {
    const apiKey = String.fromEnvironment('OPENAI_API_KEY');
    
    try {
      final formData = FormData.fromMap({
        'image': await MultipartFile.fromFile(imageFile.path),
        'prompt': _getStylePrompt(style),
        'n': 1,
        'size': '512x512',
      });

      final response = await _dio.post(
        'https://api.openai.com/v1/images/edits',
        data: formData,
        options: Options(
          headers: {
            'Authorization': 'Bearer $apiKey',
          },
        ),
      );

      return AvatarResult(
        imageUrl: response.data['data'][0]['url'],
        thumbnailUrl: response.data['data'][0]['url'],
        style: style,
      );
    } catch (e) {
      throw AvatarGenerationException('OpenAI API error: $e');
    }
  }

  /// Option 4: Stability AI API
  Future<AvatarResult> _generateWithStability(
    File imageFile,
    String style,
  ) async {
    // Similar implementation to OpenAI
    throw UnimplementedError('Stability AI integration coming soon');
  }

  Future<AvatarResult> _pollReplicateResult(
    String predictionId,
    String apiToken,
  ) async {
    for (var i = 0; i < 30; i++) {
      await Future.delayed(const Duration(seconds: 2));

      final response = await _dio.get(
        'https://api.replicate.com/v1/predictions/$predictionId',
        options: Options(
          headers: {'Authorization': 'Token $apiToken'},
        ),
      );

      final status = response.data['status'];
      if (status == 'succeeded') {
        return AvatarResult(
          imageUrl: response.data['output'][0],
          thumbnailUrl: response.data['output'][0],
          style: 'custom',
        );
      } else if (status == 'failed') {
        throw AvatarGenerationException('Generation failed');
      }
    }

    throw AvatarGenerationException('Generation timeout');
  }

  String _getReplicateModelVersion(String style) {
    // Map styles to Replicate model versions
    switch (style) {
      case 'anime':
        return 'cjwbw/anything-v3.0:...';
      case 'cartoon':
        return 'jagilley/controlnet-canny:...';
      default:
        return 'tencentarc/gfpgan:...';
    }
  }

  String _getStylePrompt(String style) {
    switch (style) {
      case 'cartoon':
        return '2D cartoon character, vibrant colors, simple shapes, friendly expression';
      case 'anime':
        return 'anime style character, manga art, cel shading, expressive eyes';
      case 'pixel':
        return 'pixel art character, 8-bit style, retro gaming aesthetic';
      case 'realistic':
        return 'realistic portrait, detailed features, natural lighting';
      default:
        return '2D character illustration';
    }
  }
}

class AvatarResult {
  final String imageUrl;
  final String thumbnailUrl;
  final String style;

  AvatarResult({
    required this.imageUrl,
    required this.thumbnailUrl,
    required this.style,
  });
}

class AvatarGenerationException implements Exception {
  final String message;
  AvatarGenerationException(this.message);

  @override
  String toString() => message;
}

