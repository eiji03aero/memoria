'use client';

import { usePathname, useRouter } from 'next/navigation';
import { useTranslation } from 'react-i18next';
import { BottomNavigation } from '@repo/design-system';
import { css } from '@/../styled-system/css';

export const Layout = ({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) => {
  const router = useRouter();
  const pathname = usePathname();
  const { t } = useTranslation();

  return (
    <div className={Styles.layout}>
      <main className={Styles.content}>{children}</main>
      <BottomNavigation>
        <BottomNavigation.Item
          active={pathname === '/timeline'}
          label={t('w.timeline')}
          icon="news"
          onPress={() => router.push('/timeline')}
        />
        <BottomNavigation.Item
          active={pathname.startsWith('/albums') || pathname.startsWith('/media')}
          label={t('w.albums')}
          icon="image-album"
          onPress={() => router.push('/albums')}
        />
        <BottomNavigation.Item
          active={pathname === '/slides'}
          label={t('w.slides')}
          icon="image-carousel"
          onPress={() => router.push('/slides')}
        />
        <BottomNavigation.Item
          active={pathname.startsWith('/account')}
          label={t('w.account')}
          icon="user"
          onPress={() => router.push('/account')}
        />
      </BottomNavigation>
    </div>
  );
};

const Styles = {
  layout: css({
    width: '100%',
    height: '100dvh',
    display: 'flex',
    flexDir: 'column',
  }),
  header: css({
    width: '100%',
    height: '3rem',
    display: 'flex',
    alignItems: 'center',
    borderBottom: '1px solid',
    borderBottomColor: 'gray.200',
    paddingX: '0.5rem',
  }),
  content: css({
    minHeight: 0,
    flex: 1,
    overflow: 'auto',
    backgroundColor: 'gray.100',
  }),
};
