'use client';

import { Provider as DesignSystemProvider } from '@repo/design-system';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { I18nextProvider } from 'react-i18next';
import i18n from '@/modules/i18n/config';
import * as Datadog from '@/modules/lib/datadog';

const queryClient = new QueryClient();
Datadog.init();

type Props = {
  children: React.ReactNode;
};

export const Providers = ({ children }: Props) => {
  return (
    <QueryClientProvider client={queryClient}>
      <DesignSystemProvider>
        <I18nextProvider i18n={i18n}></I18nextProvider>
        {children}
      </DesignSystemProvider>
    </QueryClientProvider>
  );
};
