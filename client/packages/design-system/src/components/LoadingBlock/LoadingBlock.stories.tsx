import type { Meta, StoryObj } from '@storybook/react';

import { LoadingBlock } from './LoadingBlock';

const meta: Meta<typeof LoadingBlock> = {
  component: LoadingBlock,
  tags: ['autodocs'],
};

export default meta;

type Story = StoryObj<typeof LoadingBlock>;

export const Default: Story = {};
