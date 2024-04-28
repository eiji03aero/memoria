'use client';

import {
  lightTheme,
  Provider as ReactSpectrumProvider,
} from '@adobe/react-spectrum';

type Props = {
  children: React.ReactNode;
};

export const Provider = ({ children }: Props) => {
  return (
    <ReactSpectrumProvider colorScheme="light" theme={lightTheme}>
      {children}
    </ReactSpectrumProvider>
  );
};
