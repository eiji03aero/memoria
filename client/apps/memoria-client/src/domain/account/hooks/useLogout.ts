import * as React from 'react';
import { useRouter } from 'next/navigation';

import { deleteJwt } from '@/domain/account/services';

export const useLogout = () => {
  const router = useRouter();

  const logout = React.useCallback(() => {
    deleteJwt();
    router.push('/timeline');
  }, [router]);

  return {
    logout,
  };
};
