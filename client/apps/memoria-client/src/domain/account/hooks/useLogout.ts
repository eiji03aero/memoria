import * as React from 'react';
import { useRouter } from 'next/navigation';

import { useToast } from '@/domain/common/hooks/useToast';
import { deleteJwt } from '@/domain/account/services';

export const useLogout = () => {
  const router = useRouter();
  const toast = useToast();

  const logout = React.useCallback(() => {
    deleteJwt();
    router.push('/login');
    toast.logoutSuccess();
  }, [router, toast]);

  return {
    logout,
  };
};
