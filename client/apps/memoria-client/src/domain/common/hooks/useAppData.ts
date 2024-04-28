import { proxy, useSnapshot } from 'valtio';

import * as config from '@/config';
import { axios } from '@/modules/lib/axios';

export type Response = {
  user: {
    id: string;
    name: string;
  };
  user_space: {
    id: string;
    name: string;
  };
};

export const state = proxy({
  user: { id: '', name: '' },
  userSpace: { id: '', name: '' },
});

export const restoreStateFromResponse = (data: Response) => {
  state.user = {
    id: data.user.id,
    name: data.user.name,
  };
  state.userSpace = {
    id: data.user_space.id,
    name: data.user_space.name,
  };
};

export const request = () => axios.get(`${config.ApiHost}/api/auth/app-data`);

export const useAppData = () => {
  const snapshot = useSnapshot(state);

  return {
    state,
    snapshot,
  };
};
