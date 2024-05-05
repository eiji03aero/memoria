import type { Meta, StoryObj } from '@storybook/react';

import { IconLinkCard } from './IconLinkCard';

const meta: Meta<typeof IconLinkCard> = {
  component: IconLinkCard,
  tags: ['autodocs'],
  args: {
    iconName: 'image-album',
    label: 'Image algum',
    href: '#',
  },
};

export default meta;

type Story = StoryObj<typeof IconLinkCard>;

export const Default: Story = {};
