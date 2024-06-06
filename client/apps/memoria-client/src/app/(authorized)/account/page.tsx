'use client';

import { useTranslation } from 'react-i18next';
import { css } from '@/../styled-system/css';
import { Button, IconLinkCard } from '@repo/design-system';

import { useLogout } from '@/domain/account/hooks/useLogout';
import { TopicHeader } from '@/domain/common/standalone/TopicHeader';
import { LocaleLanguageSelect } from '@/domain/common/standalone/LocalLanguageSelect';

export default function Account() {
  const { logout } = useLogout();
  const { t } = useTranslation();

  return (
    <>
      <TopicHeader label={t('w.account')} />
      <div className={Styles.section}>
        <div className={Styles.iconLinkCards}>
          <IconLinkCard
            className={Styles.iconLinkCard}
            iconName="user-add"
            label={t('w.invite-user')}
            href="/account/invite-user"
          />
        </div>
      </div>
      <div className={Styles.section}>
        <Button UNSAFE_className={Styles.button} variant="primary" onPress={logout}>
          {t('w.logout')}
        </Button>
      </div>
      <div className={Styles.section}>
        <LocaleLanguageSelect />
      </div>
    </>
  );
}

const Styles = {
  section: css({
    padding: '0.5rem',
  }),
  iconLinkCards: css({
    display: 'flex',
    flexWrap: 'wrap',
    gap: '0.5rem',
  }),
  iconLinkCard: css({
    flex: 'calc(50% - 0.5rem)',
  }),
  button: css({
    width: '100%',
  }),
};
