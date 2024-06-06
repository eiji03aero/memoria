'use client';

import { useTranslation } from 'react-i18next';
import Image from 'next/image';
import { Button } from '@repo/design-system';
import { css } from '@/../styled-system/css';

import { Link } from '@/modules/components';

export default function InternalServerError() {
  const { t } = useTranslation();
  return (
    <div
      className={css({
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        gap: '1rem',
        width: '100%',
        height: '100dvh',
        backgroundColor: 'red.400',
        padding: '1rem',
      })}
    >
      <Image
        src="/images/Logo-horizontal-black.png"
        width={300}
        height={72}
        alt="service logo"
        className={css({
          marginBottom: '1rem',
        })}
      />
      <p
        className={css({
          fontSize: '1rem',
          color: 'black',
        })}
      >
        {t('w.internal-server-error')}
      </p>
      <Button variant="primary" elementType={Link} href="/timeline">
        {t('w.back-to-data', { data: t('w.top') })}
      </Button>
    </div>
  );
}
