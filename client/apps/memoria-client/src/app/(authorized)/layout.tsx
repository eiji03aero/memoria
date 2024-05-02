import { cookies } from 'next/headers';
import { redirect } from 'next/navigation';

import * as config from '@/config';
import { Response } from '@/domain/common/hooks/useAppData';
import { AppDataBridge } from '@/domain/common/components/AppDataBridge';

import { Layout } from '@/app/(authorized)/_standalone/Layout';

export default async function AuthorizedLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const jwt = cookies().get('jwt')!.value;
  const res = await fetch(`${config.ApiHostInServer}/api/auth/app-data`, {
    cache: 'no-store',
    headers: {
      Authorization: `Bearer ${jwt}`,
    },
  });
  if (!res.ok) {
    console.error(await res.json());
    redirect('/internal-server-error');
  }

  const appData: Response = await res.json();

  return (
    <AppDataBridge appData={appData}>
      <Layout>{children}</Layout>
    </AppDataBridge>
  );
}
