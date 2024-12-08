import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:flutter_hooks/flutter_hooks.dart';

import 'package:memoria/ui/screens/welcome/welcome_screen.dart';
import 'package:memoria/ui/screens/signup/signup_screen.dart';
import 'package:memoria/ui/screens/signup_guide/signup_guide_screen.dart';
import 'package:memoria/ui/screens/login/login_screen.dart';
import 'package:memoria/ui/screens/timeline/timeline_screen.dart';

class AppTemplate extends HookWidget {
  final Widget child;
  const AppTemplate({super.key, required this.child});

  @override
  Widget build(BuildContext context) {
    final selectedIdx = useState(0);

    return Scaffold(
      body: child,
      bottomNavigationBar: BottomNavigationBar(
        items: const <BottomNavigationBarItem>[
          BottomNavigationBarItem(
            icon: Icon(Icons.newspaper),
            label: 'Timeline',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.business),
            label: 'Business',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.school),
            label: 'School',
          ),
        ],
        onTap: (idx) {
          selectedIdx.value = idx;
        },
        currentIndex: selectedIdx.value,
        selectedItemColor: Theme.of(context).primaryColor,
      ),
    );
  }
}

final goRouter = GoRouter(
  initialLocation: '/welcome',
  routes: [
    GoRoute(
      path: '/welcome',
      name: 'welcome',
      pageBuilder: (context, state) {
        return MaterialPage(key: state.pageKey, child: const WelcomeScreen());
      },
    ),
    GoRoute(
      path: '/signup',
      name: 'signup',
      pageBuilder: (context, state) {
        return MaterialPage(key: state.pageKey, child: SignupScreen());
      },
    ),
    GoRoute(
      path: '/signup-guide',
      name: 'signup-guide',
      pageBuilder: (context, state) {
        return MaterialPage(
            key: state.pageKey, child: const SignupGuideScreen());
      },
    ),
    GoRoute(
      path: '/login',
      name: 'login',
      pageBuilder: (context, state) {
        return MaterialPage(key: state.pageKey, child: LoginScreen());
      },
    ),
    ShellRoute(
      builder: (BuildContext context, GoRouterState state, Widget child) {
        return AppTemplate(child: child);
      },
      routes: <RouteBase>[
        GoRoute(
          path: '/timeline',
          name: 'timeline',
          pageBuilder: (context, state) {
            return MaterialPage(
                key: state.pageKey, child: const TimelineScreen());
          },
        ),
      ],
    ),
  ],
  errorPageBuilder: (context, state) => MaterialPage(
    key: state.pageKey,
    child: Scaffold(
      body: Center(
        child: Text(state.error.toString()),
      ),
    ),
  ),
);
