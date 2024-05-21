import { useQuery } from '@tanstack/react-query';

import * as config from '@/config';
import { axios, AxiosResponse } from '@/modules/lib/axios';
import { Album } from '@/domain/common/interfaces/api';

export type Req = {
  id: string;
};

export type Res = {
  album: Album;
};

export const request = (params: Req) =>
  axios.get<any, AxiosResponse<Res>>(`${config.ApiHost}/api/auth/albums/${params.id}`);

type Params = {
  id: string;
};

export const useAlbum = ({ id }: Params) => {
  const query = useQuery({
    queryKey: ['albums', id],
    queryFn: () => request({ id }),
  });

  return {
    refetch: query.refetch,
    album: query.data?.data.album,
  };
};
