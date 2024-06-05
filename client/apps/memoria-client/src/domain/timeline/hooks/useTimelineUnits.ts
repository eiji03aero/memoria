import { useInfiniteQuery } from '@tanstack/react-query';

import * as config from '@/config';
import { axios, AxiosResponse } from '@/modules/lib/axios';
import { TimelineUnit, CPagination, CPaginate } from '@/domain/common/interfaces/api';

type Request = CPaginate;

type Response = {
  units: Array<TimelineUnit>;
  cpagination: CPagination;
};

const request = (p: Request) =>
  axios.get<Request, AxiosResponse<Response>>(`${config.ApiHost}/api/auth/timeline`, {
    params: {
      cursor: p.cursor,
      cafter: p.cafter,
      cbefore: p.cbefore,
      cexclude: p.cexclude,
    },
  });

export const useTimelineUnits = () => {
  const query = useInfiniteQuery({
    queryKey: ['timeline'],
    queryFn: ({ pageParam }) => request(pageParam),
    initialPageParam: {
      cursor: '',
    },
    getNextPageParam: (lastPage, _, _2) => {
      return {
        cursor: lastPage.data.cpagination.next_cursor,
        cafter: 20,
        cexclude: true,
      };
    },
    getPreviousPageParam: (firstPage, _2, _) => {
      return {
        cursor: firstPage.data.cpagination.prev_cursor,
        cbefore: 20,
        cexclude: true,
      };
    },
  });

  return {
    ...query,
    units: query.data?.pages.flatMap(page => page.data.units),
  };
};
