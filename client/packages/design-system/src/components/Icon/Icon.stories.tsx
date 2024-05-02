import type { Meta, StoryFn } from '@storybook/react';

import { Icon, IconNames } from './Icon';
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
    <Flex gap="size-100" wrap="wrap">
      {IconNames.map(name => (
        <View
          key={name}
          padding="size-100"
          borderColor="dark"
          borderWidth="thin"
          borderRadius="small"
          width="size-1600"
          height="size-1600"
        >
          <Flex
            width="100%"
            height="100%"
            direction="column"
            alignItems="center"
            justifyContent="center"
          >
            <Icon
              name={name}
              color="gray-900"
              size="XL"
              marginBottom="size-200"
            />
            <CustomText size={100} color="gray-700">
              {name}
            </CustomText>
          </Flex>
        </View>
      ))}
    </Flex>
  );
};
