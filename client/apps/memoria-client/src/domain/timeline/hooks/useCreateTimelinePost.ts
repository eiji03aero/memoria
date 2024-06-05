import { useMutation, useQueryClient } from '@tanstack/react-query';

import * as config from '@/config';
import { axios, AxiosError } from '@/modules/lib/axios';

type Request = {
  content: string;
  mediumIDs?: string[];
};

export const request = (p: Request) =>
  axios.post(`${config.ApiHost}/api/auth/timeline`, {
    content: p.content,
    medium_ids: p.mediumIDs,
  });

export const useCreateTimelinePost = () => {
  const queryClient = useQueryClient();
  const { mutate, error } = useMutation({
    mutationFn: request,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['timeline'] });
    },
  });

  return {
    createPost: mutate,
    errorResponseBody: (error as AxiosError)?.response?.data,
  };
};
