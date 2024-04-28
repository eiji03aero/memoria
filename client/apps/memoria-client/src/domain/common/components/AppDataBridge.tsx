'use client';

import * as React from 'react';
import {
  restoreStateFromResponse,
  Response,
} from '@/domain/common/hooks/useAppData';

type Props = {
  appData: Response;
  children: React.ReactNode;
};

export const AppDataBridge = ({ appData, children }: Props) => {
  React.useState(() => {
    restoreStateFromResponse(appData);
  });

  return children;
};
