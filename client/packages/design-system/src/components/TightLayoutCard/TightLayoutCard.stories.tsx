import type { Meta, StoryFn } from '@storybook/react';

import { TightLayoutCard } from './TightLayoutCard';

const meta: Meta<typeof TightLayoutCard> = {
  component: TightLayoutCard,
  tags: ['autodocs'],
};

export default meta;

export const Default: StoryFn<typeof TightLayoutCard> = () => {
  return (
    <TightLayoutCard.Background>
      <TightLayoutCard>
        <div style={{ height: 100 }}>Card 1</div>
      </TightLayoutCard>
      <TightLayoutCard>
        <div style={{ height: 100 }}>Card 2</div>
      </TightLayoutCard>
      <TightLayoutCard>
        <div style={{ height: 100 }}>Card 3</div>
      </TightLayoutCard>
    </TightLayoutCard.Background>
  );
};
