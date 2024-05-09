import type { Meta, StoryObj } from '@storybook/react';
import { fn } from '@storybook/test';

import { Item } from './Item';

const meta: Meta<typeof Item> = {
  component: Item,
  tags: ['autodocs'],
  args: {
    label: 'Albums',
    icon: 'image-album',
    active: false,
    onPress: fn(),
  },
};

export default meta;

type Story = StoryObj<typeof Item>;

export const Default: Story = {};

export const Active: Story = {
  args: {
    active: true,
  },
};
