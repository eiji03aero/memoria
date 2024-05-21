import { proxy, useSnapshot } from 'valtio';

export const state = proxy({});

export const useAppUI = () => {
  const snapshot = useSnapshot(state);

  return {
    state,
    snapshot,
  };
};
