import * as React from 'react';
import { useRouter } from 'next/navigation';
import { useMutation } from '@tanstack/react-query';

import * as config from '@/config';
import { axios, AxiosError } from '@/modules/lib/axios';

import { saveJwt } from '@/domain/account/services';

type Request = {
  name: string;
  password: string;
  invitationID: string;
};

type Response = {
  token: string;
};

const request = (p: Request) =>
  axios.post<Response>(`${config.ApiHost}/api/public/invite-user-confirm`, {
    name: p.name,
    password: p.password,
    invitation_id: p.invitationID,
  });

export const useInviteUserConfirm = () => {
  const router = useRouter();

  const { mutate, error } = useMutation({
    mutationFn: request,
    onSuccess: ({ data }) => {
      router.push('/timeline');
      saveJwt(data.token);
    },
  });

  const inviteUserConfirm = React.useCallback(
    (params: Request) => {
      mutate(params);
    },
    [mutate],
  );

  return {
    inviteUserConfirm,
    errorResponseBody: (error as AxiosError)?.response?.data,
  };
};
