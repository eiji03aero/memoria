import type { Meta, StoryFn } from '@storybook/react';

import { Icon, IconNames } from './Icon';
import { Text } from '../Text';
import { Flex } from '../Flex';
import { View } from '../View';
import { CustomText } from '../CustomText';

const meta: Meta<typeof Icon> = {
  component: Icon,
  tags: ['autodocs'],
};

export default meta;

export const List: StoryFn<typeof Icon> = () => {
  return (
    <Flex gap="size-100">
      {IconNames.map(name => (
        <View
          key={name}
          padding="size-100"
          borderColor="dark"
          borderWidth="thin"
          borderRadius="small"
          width="size-2000"
          height="size-2000"
        >
          <Flex
            width="100%"
            height="100%"
            direction="column"
            alignItems="center"
            justifyContent="center"
          >
            <Icon name={name} size="XL" marginBottom="size-200" />
            <CustomText size={100} color="gray-700">
              {name}
            </CustomText>
          </Flex>
        </View>
      ))}
    </Flex>
  );
};
