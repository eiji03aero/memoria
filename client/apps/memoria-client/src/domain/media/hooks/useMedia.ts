import { useQuery } from '@tanstack/react-query';

import * as config from '@/config';
import { axios, AxiosResponse } from '@/modules/lib/axios';
import { Medium } from '@/domain/common/interfaces/api';

type Request = {
  albumID?: string;
};

type Response = {
  media: Array<Medium>;
};

const request = (params: Request) =>
  axios.get<Request, AxiosResponse<Response>>(`${config.ApiHost}/api/auth/media`, {
    params: {
      album_id: params.albumID,
    },
  });

type Params = {
  albumID?: string;
};

export const useMedia = ({ albumID }: Params) => {
  const query = useQuery({
    queryKey: ['media', { albumID }],
    queryFn: () => request({ albumID }),
  });

  return {
    refetch: query.refetch,
    media: query.data?.data.media,
    isFetching: query.isFetching,
  };
};
