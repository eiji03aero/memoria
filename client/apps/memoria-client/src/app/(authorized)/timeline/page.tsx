'use client';

import Image from 'next/image';
import { Link } from '@/modules/components';
import { css } from '@/../styled-system/css';

export default function Timeline() {
  return (
    <>
      <nav className={Styles.header}>
        <Link href="/timeline">
          <Image
            src="/images/Logo-horizontal-black.png"
            width={150}
            height={36}
            alt="service logo"
          />
        </Link>
      </nav>

      <h1>Timeline page will come here</h1>
    </>
  );
}

const Styles = {
  header: css({
    position: 'sticky',
    width: '100%',
    height: '3rem',
    display: 'flex',
    alignItems: 'center',
    backgroundColor: 'white',
    borderBottom: '1px solid',
    borderBottomColor: 'gray.200',
    paddingX: '0.5rem',
  }),
};
