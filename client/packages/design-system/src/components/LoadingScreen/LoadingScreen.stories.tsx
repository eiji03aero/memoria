import type { Meta, StoryObj } from '@storybook/react';

import { LoadingScreen } from './LoadingScreen';

const meta: Meta<typeof LoadingScreen> = {
  component: LoadingScreen,
  tags: ['autodocs'],
};

export default meta;

type Story = StoryObj<typeof LoadingScreen>;

export const Default: Story = {};
