import { useQuery } from '@tanstack/react-query';

import * as config from '@/config';
import { axios, AxiosResponse } from '@/modules/lib/axios';
import { Pagination } from '@/domain/common/interfaces/api';

type Request = {
  albumID?: string;
  mediumID: string;
};

type Response = {
  pagination: Pagination;
};

const request = (params: Request) =>
  axios.get<Request, AxiosResponse<Response>>(`${config.ApiHost}/api/auth/media/get-page`, {
    params: {
      album_id: params.albumID,
      medium_id: params.mediumID,
    },
  });

type Params = {
  albumID?: string;
  mediumID: string;
};

export const useMediumGetPage = ({ albumID, mediumID }: Params) => {
  const query = useQuery({
    queryKey: ['media', 'get-page', { albumID, mediumID }],
    queryFn: () => request({ albumID, mediumID }),
  });

  return {
    refetch: query.refetch,
    isFetched: query.isFetched,
    pagination: query.data?.data.pagination,
    isFetching: query.isFetching,
  };
};
