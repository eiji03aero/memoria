'use client';

import Image from 'next/image';
import { Flex, View, CustomText, Button } from '@repo/design-system';

import { Link } from '@/modules/components';

export default function SignupThanks() {
  return (
    <View width="100%" height="100dvh" backgroundColor="yellow-400">
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
        <CustomText size={200} color="gray-900">
          Thanks for joining memoria!
        </CustomText>
        <Button variant="primary" elementType={Link} href="/dashboard">
          Go to dashboard page
        </Button>
      </Flex>
    </View>
  );
}
