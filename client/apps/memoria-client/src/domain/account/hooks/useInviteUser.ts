import * as React from 'react';
import { useRouter } from 'next/navigation';
import { useMutation } from '@tanstack/react-query';

import * as config from '@/config';
import { axios, AxiosError } from '@/modules/lib/axios';

type Request = {
  email: string;
};

const request = (p: Request) =>
  axios.post(`${config.ApiHost}/api/auth/invite-user`, {
    email: p.email,
  });

export const useInviteUser = () => {
  const router = useRouter();

  const { mutate, error } = useMutation({
    mutationFn: request,
    onSuccess: () => {
      router.push('/account');
    },
  });

  const inviteUser = React.useCallback(
    (params: Request) => {
      mutate(params);
    },
    [mutate, router],
  );

  return {
    inviteUser,
    errorResponseBody: (error as AxiosError)?.response?.data,
  };
};
