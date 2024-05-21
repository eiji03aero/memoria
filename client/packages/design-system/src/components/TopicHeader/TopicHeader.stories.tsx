import type { Meta, StoryObj } from '@storybook/react';

import { TopicHeader } from './TopicHeader';

const meta: Meta<typeof TopicHeader> = {
  component: TopicHeader,
  tags: ['autodocs'],
  args: {
    label: 'Albums',
  },
};

export default meta;

type Story = StoryObj<typeof TopicHeader>;

export const Default: Story = {};

export const WithTrailing: Story = {
  args: {
    trailing: 'Could be buttons',
  },
};

export const WithBackButton: Story = {
  args: {
    onBack: console.log,
  },
};
