import * as React from 'react';

import { LoadingBlock } from '@repo/design-system';

export default function UnauthorizedLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return <React.Suspense fallback={<LoadingBlock />}>{children}</React.Suspense>;
}
