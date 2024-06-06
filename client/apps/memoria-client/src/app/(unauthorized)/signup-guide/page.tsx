'use client';

import { useTranslation } from 'react-i18next';
import Image from 'next/image';
import { css } from '@/../styled-system/css';

export default function SignupGuide() {
  const { t } = useTranslation();
  return (
    <div
      className={css({
        display: 'flex',
        flexDir: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        gap: '1.5rem',
        width: '100%',
        height: '100dvh',
        backgroundColor: 'sky.400',
      })}
    >
      <Image src="/images/Logo-horizontal-black.png" width={300} height={72} alt="service logo" />
      <div className={css({ width: '300px' })}>
        <p className={css({ fontSize: '1rem', color: 'black', whiteSpace: 'pre-wrap' })}>
          {t('p.signup-guide.headings')}
        </p>
      </div>
    </div>
  );
}
