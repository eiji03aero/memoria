export const isInServer = () => typeof window === 'undefined';
export const isInBrowser = () => !isInServer();
