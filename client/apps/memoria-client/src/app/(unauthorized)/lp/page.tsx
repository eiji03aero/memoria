'use client';

import Image from 'next/image';
import {
  Flex,
  View,
  CustomText,
  Heading,
  Content,
  Button,
  Divider,
  styles,
} from '@repo/design-system';

import { Link } from '@/modules/components';

export default function LP() {
  return (
    <>
      <View width="100%" height="size-4600" backgroundColor="yellow-400">
        <Flex
          width="100%"
          height="100%"
          direction="column"
          alignItems="center"
          justifyContent="center"
          gap="size-300"
        >
          <Image
            src="/images/Logo-horizontal-black.png"
            width={300}
            height={72}
            alt="service logo"
          />
          <CustomText size={200} color="gray-900">
            Cherish your moments for memories
          </CustomText>
          <Button variant="primary" elementType={Link} href="/signup">
            Sign up now
          </Button>
        </Flex>
      </View>

      <Content>
        <View padding="size-300">
          <Heading level={1} marginBottom="size-100">
            <CustomText size={300} color="gray-900">
              What is this?
            </CustomText>
          </Heading>
          <p className={Styles.p}>
            This app is ... well, you can say it is pretty much copy of photo
            app on ios.
            <br />
            I'm building this app for following reasons:
          </p>
          <ul className={Styles.ul}>
            <li>Have a development experience in full stack way.</li>
            <li>
              To keep myself motivated, build something that me and my family
              can actually use.
            </li>
            <li>To make my wife and baby proud.</li>
          </ul>
        </View>
      </Content>

      <Divider />

      <Content>
        <View padding="size-300">
          <Heading level={1} marginBottom="size-100">
            <CustomText size={300} color="gray-900">
              What can you do on this app?
            </CustomText>
          </Heading>
          <p className={Styles.p}>
            Glad you asked!
            <br />
            There are plenty of cool stuff you can do with this app.
          </p>

          <View marginBottom="size-100"></View>

          <Heading level={2} marginBottom="size-100">
            <CustomText size={200} color="gray-900">
              Upload to albums
            </CustomText>
          </Heading>
          <p className={Styles.p}>
            This app helps you organize your media files so that you know where
            to find them.
          </p>

          <Heading level={2} marginBottom="size-100">
            <CustomText size={200} color="gray-900">
              Share with your family, friends
            </CustomText>
          </Heading>
          <p className={Styles.p}>
            The albums will be shared among the users that belong to same user
            space.
          </p>

          <Heading level={2} marginBottom="size-100">
            <CustomText size={200} color="gray-900">
              Make comments to hold more of moments
            </CustomText>
          </Heading>
          <p className={Styles.p}>
            You can create thread on any media, so that you can cherish more of
            your memories.
          </p>
        </View>
      </Content>
    </>
  );
}

const Styles = {
  p: styles.classnames('mb-8'),
  ul: styles.classnames('list-disc', 'ml-8'),
};
