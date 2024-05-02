'use client';

import { usePathname, useRouter } from 'next/navigation';
import Image from 'next/image';
import { View, Flex, BottomNavigation } from '@repo/design-system';

import { Link } from '@/modules/components';

export const Layout = ({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) => {
  const router = useRouter();
  const pathname = usePathname();

  return (
    <Flex width="100%" height="100dvh" direction="column">
      <View
        height="size-500"
        borderBottomWidth="thin"
        borderBottomColor="dark"
        paddingX="size-100"
      >
        <Flex height="100%" alignItems="center">
          <Link href="/dashboard">
            <Image
              src="/images/Logo-horizontal-black.png"
              width={150}
              height={36}
              alt="service logo"
            />
          </Link>
        </Flex>
      </View>
      <View minHeight="size-0" flex={1} UNSAFE_className="overflow-auto">
        {children}
      </View>
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
          active={pathname === '/settings'}
          label="Settings"
          icon="settings"
          onPress={() => router.push('/settings')}
        />
      </BottomNavigation>
    </Flex>
  );
};
