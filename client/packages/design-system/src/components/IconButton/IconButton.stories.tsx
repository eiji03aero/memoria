import type { Meta, StoryObj } from '@storybook/react';

import { IconButton } from './IconButton';

const meta: Meta<typeof IconButton> = {
  component: IconButton,
  tags: ['autodocs'],
  args: {
    iconName: 'settings',
  },
};

export default meta;

type Story = StoryObj<typeof IconButton>;

export const Default: Story = {};
