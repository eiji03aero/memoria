import * as React from 'react';

import { LoadingScreen } from '@repo/design-system';

export default function UnauthorizedLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <React.Suspense fallback={<LoadingScreen />}>{children}</React.Suspense>
  );
}
