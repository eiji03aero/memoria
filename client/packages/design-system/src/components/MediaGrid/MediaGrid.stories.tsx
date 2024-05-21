import type { Meta, StoryFn } from '@storybook/react';

import { MediaGrid } from './MediaGrid';

const meta: Meta<typeof MediaGrid> = {
  component: MediaGrid,
  tags: ['autodocs'],
};

export default meta;

type Story = StoryFn<typeof MediaGrid>;

export const Default: Story = () => {
  return (
    <MediaGrid>
      <MediaGrid.Tile
        selected
        src="https://memoria-dev.s3.ap-northeast-1.amazonaws.com/media/01HXW21TD6RDR0AP8CTP1X5N09/original.png"
      />
      <MediaGrid.Tile
        selected
        src="https://memoria-dev.s3.ap-northeast-1.amazonaws.com/media/01HXW21TD6RDR0AP8CTP1X5N09/original.png"
      />
      <MediaGrid.Tile src="https://memoria-dev.s3.ap-northeast-1.amazonaws.com/media/01HXW21TD6RDR0AP8CTP1X5N09/original.png" />
      <MediaGrid.Tile src="https://memoria-dev.s3.ap-northeast-1.amazonaws.com/media/01HXW21TD6RDR0AP8CTP1X5N09/original.png" />
      <MediaGrid.Tile src="https://memoria-dev.s3.ap-northeast-1.amazonaws.com/media/01HXW21TD6RDR0AP8CTP1X5N09/original.png" />
      <MediaGrid.Tile src="https://memoria-dev.s3.ap-northeast-1.amazonaws.com/media/01HXW21TD6RDR0AP8CTP1X5N09/original.png" />
      <MediaGrid.Tile />
      <MediaGrid.Tile src="https://memoria-dev.s3.ap-northeast-1.amazonaws.com/media/01HXW21TD6RDR0AP8CTP1X5N09/original.png" />
    </MediaGrid>
  );
};
