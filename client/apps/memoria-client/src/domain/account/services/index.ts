import { Cookies } from '@/modules/lib/cookie';

export const saveJwt = (jwt: string) => {
  Cookies.set('jwt', jwt, { path: '/' });
};

export const getJwt = () => {
  return Cookies.get('jwt');
};
