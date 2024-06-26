import * as React from 'react';
import { useRouter } from 'next/navigation';
import { useMutation } from '@tanstack/react-query';

import * as config from '@/config';
import { axios, AxiosError } from '@/modules/lib/axios';

import { useToast } from '@/domain/common/hooks/useToast';
import { saveJwt } from '@/domain/account/services';

type Request = {
  email: string;
  password: string;
};

type Response = {
  token: string;
};

const request = (p: Request) =>
  axios.post<Response>(`${config.ApiHost}/api/public/login`, {
    email: p.email,
    password: p.password,
  });

export const useLogin = () => {
  const router = useRouter();
  const toast = useToast();

  const { mutate, error } = useMutation({
    mutationFn: request,
    onSuccess: ({ data }) => {
      saveJwt(data.token);
      router.push('/timeline');
      toast.loginSuccess();
    },
  });

  return {
    login: mutate,
    errorResponseBody: (error as AxiosError)?.response?.data,
  };
};
