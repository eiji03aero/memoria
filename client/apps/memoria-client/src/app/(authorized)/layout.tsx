import { cookies } from 'next/headers';

import * as config from '@/config';
import { Response } from '@/domain/common/hooks/useAppData';
import { AppDataBridge } from '@/domain/common/components/AppDataBridge';

export default async function AuthorizedLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const jwt = cookies().get('jwt')!.value;
  const appData: Response = await fetch(
    `${config.ApiHostInServer}/api/auth/app-data`,
    {
      cache: 'no-store',
      headers: {
        Authorization: `Bearer ${jwt}`,
      },
    },
  ).then(res => res.json());

  return <AppDataBridge appData={appData}>{children}</AppDataBridge>;
}
