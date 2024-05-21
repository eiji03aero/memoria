import { useMutation, useQueryClient } from '@tanstack/react-query';

import * as config from '@/config';
import { axios } from '@/modules/lib/axios';

const request = (id: string) => axios.delete(`${config.ApiHost}/api/auth/albums/${id}`);

export const useDeleteAlbum = () => {
  const queryClient = useQueryClient();
  const { mutate } = useMutation({
    mutationFn: request,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['albums'] });
    },
  });

  return {
    deleteAlbum: mutate,
  };
};
