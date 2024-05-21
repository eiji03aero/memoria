import * as React from 'react';
import { useQuery } from '@tanstack/react-query';

import * as config from '@/config';
import { axios, AxiosResponse } from '@/modules/lib/axios';
import { Album } from '@/domain/common/interfaces/api';

export type Res = {
  albums: Array<Album>;
};

export const request = () =>
  axios.get<any, AxiosResponse<Res>>(`${config.ApiHost}/api/auth/albums`);

export const useAlbums = () => {
  const query = useQuery({
    queryKey: ['albums'],
    queryFn: request,
  });

  const refetch = React.useCallback(() => {
    query.refetch();
  }, [query.refetch]);

  return {
    refetch,
    albums: query.data?.data.albums,
    isFetching: query.isFetching,
  };
};
