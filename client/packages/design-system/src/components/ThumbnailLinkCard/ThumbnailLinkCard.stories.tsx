import type { Meta, StoryObj } from '@storybook/react';

import { ThumbnailLinkCard } from './ThumbnailLinkCard';

const meta: Meta<typeof ThumbnailLinkCard> = {
  component: ThumbnailLinkCard,
  tags: ['autodocs'],
  args: {
    label: 'Baby memory',
  },
};

export default meta;

type Story = StoryObj<typeof ThumbnailLinkCard>;

export const Default: Story = {};
export const WithMore: Story = {
  args: {
    onOpenMenu: console.log,
  },
};
