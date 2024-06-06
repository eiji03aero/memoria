'use client';

import { useTranslation } from 'react-i18next';
import Image from 'next/image';
import { Heading, Content, Button, Divider, styles } from '@repo/design-system';
import { css } from '@/../styled-system/css';

import { Link } from '@/modules/components';

export default function LP() {
  const { t } = useTranslation();
  return (
    <>
      <section
        className={css({
          display: 'flex',
          flexDir: 'column',
          alignItems: 'center',
          justifyContent: 'center',
          gap: '1.5rem',
          width: '100%',
          height: '400px',
          backgroundColor: 'yellow.400',
        })}
      >
        <Image src="/images/Logo-horizontal-black.png" width={300} height={72} alt="service logo" />
        <p className={css({ fontSize: '1rem', color: 'black' })}>{t('p.lp.headings')}</p>
        <div
          className={css({
            display: 'flex',
            gap: '0.5rem',
            alignItems: 'center',
          })}
        >
          <Button variant="primary" elementType={Link} href="/signup">
            {t('w.sign-up-now')}
          </Button>
          <Button variant="primary" elementType={Link} href="/login">
            {t('w.login')}
          </Button>
        </div>
      </section>

      <Content>
        <div className={css({ padding: '1.5rem' })}>
          <Heading level={1} marginBottom="size-100">
            <span className={css({ fontSize: '1rem', color: 'black' })}>What is this?</span>
          </Heading>
          <p className={Styles.p}>
            This app is ... well, you can say it is pretty much copy of photo app on ios.
            <br />
            I'm building this app for following reasons:
          </p>
          <ul className={Styles.ul}>
            <li>Have a development experience in full stack way.</li>
            <li>
              To keep myself motivated, build something that me and my family can actually use.
            </li>
            <li>To make my wife and baby proud.</li>
          </ul>
        </div>
      </Content>

      <Divider />

      <Content>
        <div className={css({ padding: '1.5rem' })}>
          <Heading level={1} marginBottom="size-100">
            <span className={Styles.sectionTitle}>What can you do on this app?</span>
          </Heading>
          <p className={Styles.p}>
            Glad you asked!
            <br />
            There are plenty of cool stuff you can do with this app.
          </p>

          <div className={css({ marginBottom: '0.5rem' })}></div>

          <Heading level={2} marginBottom="size-100">
            <span className={Styles.sectionSubTitle}>Upload to albums</span>
          </Heading>
          <p className={Styles.p}>
            This app helps you organize your media files so that you know where to find them.
          </p>

          <Heading level={2} marginBottom="size-100">
            <span className={Styles.sectionSubTitle}>Share with your family, friends</span>
          </Heading>
          <p className={Styles.p}>
            The albums will be shared among the users that belong to same user space.
          </p>

          <Heading level={2} marginBottom="size-100">
            <span className={Styles.sectionSubTitle}>Make comments to hold more of moments</span>
          </Heading>
          <p className={Styles.p}>
            You can create thread on any media, so that you can cherish more of your memories.
          </p>
        </div>
      </Content>
    </>
  );
}

const Styles = {
  p: styles.classnames('mb-8'),
  ul: styles.classnames(
    'list-disc',
    'ml-8',
    css({
      listStyle: 'unset',
    }),
  ),
  sectionTitle: css({
    fontSize: '1.5rem',
    color: 'black',
  }),
  sectionSubTitle: css({
    fontSize: '1rem',
    color: 'black',
  }),
};
