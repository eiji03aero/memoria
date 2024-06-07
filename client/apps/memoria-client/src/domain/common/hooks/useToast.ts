import * as React from 'react';
import { useTranslation } from 'react-i18next';
import { ToastQueue } from '@repo/design-system';

export const useToast = () => {
  const { t } = useTranslation();

  const opt = React.useMemo(
    () => ({
      timeout: 8000,
    }),
    [],
  );

  const changedLocaleLanguage = React.useCallback(() => {
    ToastQueue.positive(t('s.changed-locale-language'), opt);
  }, [t, opt]);

  const inviteUserSuccess = React.useCallback(() => {
    ToastQueue.positive(t('s.invite-user-success'), opt);
  }, [t, opt]);

  const logoutSuccess = React.useCallback(() => {
    ToastQueue.positive(t('s.logout-success'), opt);
  }, [t, opt]);

  const loginSuccess = React.useCallback(() => {
    ToastQueue.positive(t('s.login-success'), opt);
  }, [t, opt]);

  return React.useMemo(
    () => ({
      changedLocaleLanguage,
      inviteUserSuccess,
      logoutSuccess,
      loginSuccess,
    }),
    [changedLocaleLanguage, inviteUserSuccess, logoutSuccess],
  );
};
