'use client';

import Image from 'next/image';
import { Flex, View, CustomText } from '@repo/design-system';

export default function SignupGuide() {
  return (
    <View width="100%" height="100dvh" backgroundColor="seafoam-400">
      <Flex
        width="100%"
        height="100%"
        direction="column"
        justifyContent="center"
        alignItems="center"
        gap="size-300"
      >
        <Image
          src="/images/Logo-horizontal-black.png"
          width={300}
          height={72}
          alt="service logo"
        />
        <View maxWidth="size-3000">
          <CustomText size={200} color="gray-900">
            Thanks for signing up memoria!
            <br />
            We have sent you an email to verify your email address.
            <br />
            Please check it to complete signup process when you got time :)
          </CustomText>
        </View>
      </Flex>
    </View>
  );
}
