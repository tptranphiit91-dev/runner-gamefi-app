import 'dart:async';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:geolocator/geolocator.dart';
import 'package:go_router/go_router.dart';
import '../providers/run_provider.dart';

class RunScreen extends ConsumerStatefulWidget {
  const RunScreen({super.key});

  @override
  ConsumerState<RunScreen> createState() => _RunScreenState();
}

class _RunScreenState extends ConsumerState<RunScreen> {
  bool _isRunning = false;
  bool _isSubmitting = false;
  int _elapsedSeconds = 0;
  double _totalDistance = 0.0; // in meters
  Timer? _timer;
  Position? _lastPosition;
  DateTime? _startTime;
  StreamSubscription<Position>? _positionStream;

  @override
  void dispose() {
    _timer?.cancel();
    _positionStream?.cancel();
    super.dispose();
  }

  Future<void> _startRun() async {
    // Check location permissions
    bool serviceEnabled = await Geolocator.isLocationServiceEnabled();
    if (!serviceEnabled) {
      _showError('Location services are disabled');
      return;
    }

    LocationPermission permission = await Geolocator.checkPermission();
    if (permission == LocationPermission.denied) {
      permission = await Geolocator.requestPermission();
      if (permission == LocationPermission.denied) {
        _showError('Location permissions are denied');
        return;
      }
    }

    if (permission == LocationPermission.deniedForever) {
      _showError('Location permissions are permanently denied');
      return;
    }

    // Start tracking
    setState(() {
      _isRunning = true;
      _elapsedSeconds = 0;
      _totalDistance = 0.0;
      _startTime = DateTime.now();
    });

    // Start timer
    _timer = Timer.periodic(const Duration(seconds: 1), (timer) {
      setState(() {
        _elapsedSeconds++;
      });
    });

    // Start position tracking
    const LocationSettings locationSettings = LocationSettings(
      accuracy: LocationAccuracy.high,
      distanceFilter: 10, // Update every 10 meters
    );

    _positionStream = Geolocator.getPositionStream(
      locationSettings: locationSettings,
    ).listen((Position position) {
      if (_lastPosition != null) {
        double distance = Geolocator.distanceBetween(
          _lastPosition!.latitude,
          _lastPosition!.longitude,
          position.latitude,
          position.longitude,
        );
        setState(() {
          _totalDistance += distance;
        });
      }
      _lastPosition = position;
    });
  }

  Future<void> _stopRun() async {
    if (!_isRunning) return;

    setState(() {
      _isRunning = false;
      _isSubmitting = true;
    });

    _timer?.cancel();
    await _positionStream?.cancel();

    // Submit run to backend
    try {
      final endTime = DateTime.now();
      await ref.read(runProvider.notifier).submitRun(
            distance: _totalDistance.round(),
            duration: _elapsedSeconds,
            startedAt: _startTime!.toIso8601String(),
            endedAt: endTime.toIso8601String(),
          );

      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text(
              'Run submitted! Distance: ${(_totalDistance / 1000).toStringAsFixed(2)} km',
            ),
            backgroundColor: Colors.green,
          ),
        );
        
        // Reset state
        setState(() {
          _elapsedSeconds = 0;
          _totalDistance = 0.0;
          _lastPosition = null;
          _startTime = null;
        });
      }
    } catch (e) {
      _showError(e.toString().replaceAll('Exception: ', ''));
    } finally {
      if (mounted) {
        setState(() {
          _isSubmitting = false;
        });
      }
    }
  }

  void _showError(String message) {
    ScaffoldMessenger.of(context).showSnackBar(
      SnackBar(
        content: Text(message),
        backgroundColor: Colors.red,
      ),
    );
  }

  String _formatTime(int seconds) {
    final hours = seconds ~/ 3600;
    final minutes = (seconds % 3600) ~/ 60;
    final secs = seconds % 60;
    return '${hours.toString().padLeft(2, '0')}:${minutes.toString().padLeft(2, '0')}:${secs.toString().padLeft(2, '0')}';
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Run Tracker'),
        leading: IconButton(
          icon: const Icon(Icons.arrow_back),
          onPressed: () => context.go('/home'),
        ),
      ),
      body: Center(
        child: Padding(
          padding: const EdgeInsets.all(24.0),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              // Timer Display
              Container(
                padding: const EdgeInsets.all(32),
                decoration: BoxDecoration(
                  color: Colors.blue.shade50,
                  shape: BoxShape.circle,
                ),
                child: Column(
                  children: [
                    const Icon(
                      Icons.timer,
                      size: 48,
                      color: Colors.blue,
                    ),
                    const SizedBox(height: 16),
                    Text(
                      _formatTime(_elapsedSeconds),
                      style: const TextStyle(
                        fontSize: 48,
                        fontWeight: FontWeight.bold,
                        fontFamily: 'monospace',
                      ),
                    ),
                  ],
                ),
              ),
              const SizedBox(height: 48),

              // Distance Display
              Card(
                child: Padding(
                  padding: const EdgeInsets.all(24.0),
                  child: Column(
                    children: [
                      const Text(
                        'Distance',
                        style: TextStyle(
                          fontSize: 18,
                          color: Colors.grey,
                        ),
                      ),
                      const SizedBox(height: 8),
                      Text(
                        '${(_totalDistance / 1000).toStringAsFixed(2)} km',
                        style: const TextStyle(
                          fontSize: 36,
                          fontWeight: FontWeight.bold,
                        ),
                      ),
                    ],
                  ),
                ),
              ),
              const SizedBox(height: 48),

              // Start/Stop Button
              SizedBox(
                width: double.infinity,
                height: 60,
                child: FilledButton(
                  onPressed: _isSubmitting
                      ? null
                      : (_isRunning ? _stopRun : _startRun),
                  style: FilledButton.styleFrom(
                    backgroundColor: _isRunning ? Colors.red : Colors.green,
                  ),
                  child: _isSubmitting
                      ? const CircularProgressIndicator(color: Colors.white)
                      : Text(
                          _isRunning ? 'Stop Run' : 'Start Run',
                          style: const TextStyle(fontSize: 20),
                        ),
                ),
              ),
            ],
          ),
        ),
      ),
      bottomNavigationBar: BottomNavigationBar(
        currentIndex: 1,
        onTap: (index) {
          if (_isRunning) {
            _showError('Please stop the run before navigating');
            return;
          }
          switch (index) {
            case 0:
              context.go('/home');
              break;
            case 1:
              // Already on run screen
              break;
            case 2:
              context.go('/shop');
              break;
          }
        },
        items: const [
          BottomNavigationBarItem(
            icon: Icon(Icons.home),
            label: 'Home',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.directions_run),
            label: 'Run',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.shopping_bag),
            label: 'Shop',
          ),
        ],
      ),
    );
  }
}

