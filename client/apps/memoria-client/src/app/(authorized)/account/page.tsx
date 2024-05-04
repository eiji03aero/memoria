'use client';

import { Button } from '@repo/design-system';

import { useLogout } from '@/domain/account/hooks/useLogout';

export default function Account() {
  const { logout } = useLogout();

  return (
    <>
      <Button variant="primary" onPress={logout}>
        Logout
      </Button>
    </>
  );
}
