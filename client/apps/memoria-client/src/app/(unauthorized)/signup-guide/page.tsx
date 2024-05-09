'use client';

import Image from 'next/image';
import { css } from '@/../styled-system/css';

export default function SignupGuide() {
  return (
    <div
      className={css({
        display: 'flex',
        flexDir: 'column',
        justifyContent: 'center',
        alignItems: 'cetner',
        gap: '1.5rem',
        width: '100%',
        height: '100dvh',
        backgroundColor: 'sky.400',
      })}
    >
      <Image
        src="/images/Logo-horizontal-black.png"
        width={300}
        height={72}
        alt="service logo"
      />
      <div className={css({ width: '300px' })}>
        <p className={css({ fontSize: '1rem', color: 'black' })}>
          Thanks for signing up memoria!
          <br />
          We have sent you an email to verify your email address.
          <br />
          Please check it to complete signup process when you got time :)
        </p>
      </div>
    </div>
  );
}
