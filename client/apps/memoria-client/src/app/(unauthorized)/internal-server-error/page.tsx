'use client';

import Image from 'next/image';
import { Flex, View, CustomText, Button } from '@repo/design-system';

import { Link } from '@/modules/components';

export default function InternalServerError() {
  return (
    <View
      width="100%"
      height="100dvh"
      backgroundColor="red-400"
      padding="size-100"
    >
      <Flex
        width="100%"
        height="100dvh"
        alignItems="center"
        justifyContent="center"
        direction="column"
        gap="size-200"
      >
        <Image
          src="/images/Logo-horizontal-black.png"
          width={300}
          height={72}
          alt="service logo"
        />
        <CustomText size={200} color="gray-900">
          Internal server error
        </CustomText>
        <Button variant="primary" elementType={Link} href="/dashboard">
          Back to dashboard
        </Button>
      </Flex>
    </View>
  );
}
