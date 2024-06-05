'use client';

import { css } from '@/../styled-system/css';
import { Button, IconLinkCard } from '@repo/design-system';

import { useLogout } from '@/domain/account/hooks/useLogout';
import { TopicHeader } from '@/domain/common/standalone/TopicHeader';

export default function Account() {
  const { logout } = useLogout();

  return (
    <>
      <TopicHeader label="Account" />
      <div className={Styles.section}>
        <div className={Styles.iconLinkCards}>
          <IconLinkCard
            className={Styles.iconLinkCard}
            iconName="user-add"
            label="Invite user"
            href="/account/invite-user"
          />
          <IconLinkCard
            className={Styles.iconLinkCard}
            iconName="image-album"
            label="Image album"
            href="#"
          />
        </div>
      </div>
      <div className={Styles.section}>
        <Button UNSAFE_className={Styles.button} variant="primary" onPress={logout}>
          Logout
        </Button>
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
