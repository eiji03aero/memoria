import axios, { AxiosError } from 'axios';

import { getJwt } from '@/domain/account/services';

axios.interceptors.request.use(
  function (config) {
    const jwt = getJwt();
    if (jwt) {
      config.headers.Authorization = `Bearer ${jwt}`;
    }

    return config;
  },
  function (error) {
    // Do something with request error
    return Promise.reject(error);
  },
);

export { axios, AxiosError };
