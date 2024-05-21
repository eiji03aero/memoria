import { useMutation, useQueryClient } from '@tanstack/react-query';

import * as config from '@/config';
import { axios } from '@/modules/lib/axios';

type Request = {
  albumIDs: string[];
  mediumIDs: string[];
};

export const request = (p: Request) =>
  axios.post(`${config.ApiHost}/api/auth/albums/remove-media`, {
    album_ids: p.albumIDs,
    medium_ids: p.mediumIDs,
  });

export const useRemoveMediaFromAlbums = () => {
  const queryClient = useQueryClient();
  const { mutate } = useMutation({
    mutationFn: request,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['media'] });
    },
  });

  return {
    removeMediaFromAlbums: mutate,
  };
};
