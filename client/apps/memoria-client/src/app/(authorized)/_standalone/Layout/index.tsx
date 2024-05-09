'use client';

import { usePathname, useRouter } from 'next/navigation';
import Image from 'next/image';
import { BottomNavigation } from '@repo/design-system';
import { css } from '@/../styled-system/css';

import { Link } from '@/modules/components';

export const Layout = ({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) => {
  const router = useRouter();
  const pathname = usePathname();

  return (
    <div className={Styles.layout}>
      <nav className={Styles.header}>
        <Link href="/dashboard">
          <Image
            src="/images/Logo-horizontal-black.png"
            width={150}
            height={36}
            alt="service logo"
          />
        </Link>
      </nav>
      <main className={Styles.content}>{children}</main>
      <BottomNavigation>
        <BottomNavigation.Item
          active={pathname === '/dashboard'}
          label="Dashboard"
          icon="article"
          onPress={() => router.push('/dashboard')}
        />
        <BottomNavigation.Item
          active={pathname === '/albums'}
          label="Albums"
          icon="image-album"
          onPress={() => router.push('/albums')}
        />
        <BottomNavigation.Item
          active={pathname === '/slides'}
          label="Slides"
          icon="image-carousel"
          onPress={() => router.push('/slides')}
        />
        <BottomNavigation.Item
          active={pathname.startsWith('/account')}
          label="Account"
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
