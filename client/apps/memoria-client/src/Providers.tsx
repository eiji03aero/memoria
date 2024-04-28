'use client';

import { Provider as DesignSystemProvider } from '@repo/design-system';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

const queryClient = new QueryClient();

type Props = {
  children: React.ReactNode;
};

export const Providers = ({ children }: Props) => {
  return (
    <QueryClientProvider client={queryClient}>
      <DesignSystemProvider>{children}</DesignSystemProvider>
    </QueryClientProvider>
  );
};
