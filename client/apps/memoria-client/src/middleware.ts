import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';

const PublicRoutes = [
  '/lp',
  '/signup',
  '/signup-guide',
  '/signup-thanks',
  '/login',
  '/invite-user-confirm',
  '/internal-server-error',
];

const isAuthenticated = (req: NextRequest) => {
  return req.cookies.has('jwt');
};

const isPublicRoutes = (path: string) => {
  return PublicRoutes.includes(path);
};

export default function middleware(req: NextRequest) {
  if (!isAuthenticated(req) && !isPublicRoutes(req.nextUrl.pathname)) {
    const absoluteURL = new URL('/lp', req.nextUrl.origin);
    return NextResponse.redirect(absoluteURL.toString());
  }
}

export const config = {
  matcher: [
    /*
     * Match all request paths except for the ones starting with:
     * - api (API routes)
     * - _next/static (static files)
     * - _next/image (image optimization files)
     * - favicon.ico (favicon file)
     */
    '/((?!api|_next/static|_next/image|images|favicon.ico).*)',
  ],
};
