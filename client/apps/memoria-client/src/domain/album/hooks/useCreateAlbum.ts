import { useMutation, useQueryClient } from '@tanstack/react-query';

import * as config from '@/config';
import { axios, AxiosError } from '@/modules/lib/axios';

type Request = {
  name: string;
};

export const request = (p: Request) =>
  axios.post(`${config.ApiHost}/api/auth/albums`, {
    name: p.name,
  });

export const useCreateAlbum = () => {
  const queryClient = useQueryClient();
  const { mutate, error } = useMutation({
    mutationFn: request,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['albums'] });
    },
  });

  return {
    createAlbum: mutate,
    errorResponseBody: (error as AxiosError)?.response?.data,
  };
};
