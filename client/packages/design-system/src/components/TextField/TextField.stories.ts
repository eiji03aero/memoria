import type { Meta, StoryObj } from '@storybook/react';

import { TextField } from './TextField';

// More on how to set up stories at: https://storybook.js.org/docs/writing-stories#default-export
const meta: Meta<typeof TextField> = {
  component: TextField,
  tags: ['autodocs'],
  args: {
    label: 'Albums',
  },
};

export default meta;

type Story = StoryObj<typeof TextField>;

export const ErrorMessage: Story = {
  args: {
    validationState: 'invalid',
    errorMessage: 'You are not entitled',
  },
};
