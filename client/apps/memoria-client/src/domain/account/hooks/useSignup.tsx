import * as React from 'react';
import { useRouter } from 'next/navigation';
import { useMutation } from '@tanstack/react-query';

import * as config from '@/config';
import { axios, AxiosError } from '@/modules/lib/axios';

import { saveJwt } from '@/domain/account/services';

type Request = {
  name: string;
  email: string;
  userSpaceName: string;
  password: string;
};

type Response = {
  token: string;
};

const request = (p: Request) =>
  axios.post<Response>(`${config.ApiHost}/api/public/signup`, {
    name: p.name,
    email: p.email,
    user_space_name: p.userSpaceName,
    password: p.password,
  });

export const useSignup = () => {
  const router = useRouter();

  const { mutate, error } = useMutation({
    mutationFn: request,
    onSuccess: ({ data }) => {
      router.push('/signup-guide');
      saveJwt(data.token);
    },
  });

  const signup = React.useCallback(
    (params: Request) => {
      mutate(params);
    },
    [mutate, router],
  );

  return {
    signup,
    errorResponseBody: (error as AxiosError)?.response?.data,
  };
};
