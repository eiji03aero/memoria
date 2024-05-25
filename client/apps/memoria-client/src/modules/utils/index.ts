export const isInServer = () => typeof window === 'undefined';
export const isInBrowser = () => !isInServer();

export const buildQueryParams = (params: {
  [key: string]: string | number | Array<any> | null | undefined;
}) => {
  const p = new URLSearchParams();

  Object.entries(params).map(([key, value]) => {
    if (!value) {
      return;
    }

    p.append(key, value.toString());
  });

  return p.toString();
};
