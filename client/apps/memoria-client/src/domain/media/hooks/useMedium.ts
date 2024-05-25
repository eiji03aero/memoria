import { useQuery } from '@tanstack/react-query';

import * as config from '@/config';
import { axios, AxiosResponse } from '@/modules/lib/axios';
import { Medium } from '@/domain/common/interfaces/api';

type Request = {
  mediumID: string;
};

type Response = {
  medium: Medium;
};

const request = (params: Request) =>
  axios.get<Request, AxiosResponse<Response>>(
    `${config.ApiHost}/api/auth/media/${params.mediumID}`,
  );

type Params = {
  mediumID: string;
};

export const useMedium = ({ mediumID }: Params) => {
  const query = useQuery({
    queryKey: ['media', { mediumID }],
    queryFn: () => request({ mediumID }),
  });

  return {
    refetch: query.refetch,
    medium: query.data?.data.medium,
    isFetching: query.isFetching,
  };
};
