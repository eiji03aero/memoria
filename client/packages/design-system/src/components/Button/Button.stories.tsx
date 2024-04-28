import type { Meta } from '@storybook/react';

import { Button } from './index';

const meta: Meta<typeof Button> = {
  component: Button,
  tags: ['autodocs'],
  args: {
    children: 'Button',
  },
  argTypes: {
    variant: {
      options: [
        'primary',
        'secondary',
        'negative',
        'cta',
        'accent',
        'overBackground',
      ],
      control: 'select',
    },
  },
};

export default meta;

export const Default = {};
