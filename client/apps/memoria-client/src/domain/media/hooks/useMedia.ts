import { useInfiniteQuery } from '@tanstack/react-query';

import * as config from '@/config';
import { axios, AxiosResponse } from '@/modules/lib/axios';
import { Medium, Pagination, Paginate } from '@/domain/common/interfaces/api';

type Request = Paginate & {
  albumID?: string;
};

type Response = {
  media: Array<Medium>;
  pagination: Pagination | undefined;
};

const request = (p: Request) =>
  axios.get<Request, AxiosResponse<Response>>(`${config.ApiHost}/api/auth/media`, {
    params: {
      album_id: p.albumID,
      page: p.page,
      per_page: p.perPage,
    },
  });

type Params = {
  albumID?: string;
  enabled?: boolean;
  initialPosition?: number;
  perPage?: number;
};

export const useMedia = ({
  albumID,
  enabled = true,
  initialPosition = 1,
  perPage = 100,
}: Params) => {
  const query = useInfiniteQuery({
    queryKey: ['media', { albumID }],
    queryFn: ({ pageParam }) => request({ albumID, page: pageParam, perPage }),
    enabled,
    initialPageParam: Math.ceil(initialPosition / perPage),
    getNextPageParam: (lastPage, _, lastPageParam) => {
      if (lastPage.data.media.length === 0) {
        return undefined;
      }
      return lastPageParam + 1;
    },
    getPreviousPageParam: (_, _2, firstPageParam) => {
      if (firstPageParam <= 1) {
        return undefined;
      }
      return firstPageParam - 1;
    },
  });

  return {
    ...query,
    media: query.data?.pages.flatMap(page => page.data.media),
  };
};
