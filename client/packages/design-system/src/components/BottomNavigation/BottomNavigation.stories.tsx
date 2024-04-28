import * as React from 'react';
import type { Meta, StoryFn } from '@storybook/react';

import { BottomNavigation } from './BottomNavigation';

const meta: Meta<typeof BottomNavigation> = {
  component: BottomNavigation,
  tags: ['autodocs'],
  args: {},
};

export default meta;

// More on writing stories with args: https://storybook.js.org/docs/writing-stories/args
export const Default: StoryFn<typeof BottomNavigation> = () => {
  const [activeMenu, setActiveMenu] = React.useState('albums');

  return (
    <BottomNavigation>
      <BottomNavigation.Item
        active={activeMenu === 'albums'}
        label="Albums"
        icon="image-album"
        onPress={() => setActiveMenu('albums')}
      />
      <BottomNavigation.Item
        active={activeMenu === 'add'}
        label="Add"
        icon="add"
        onPress={() => setActiveMenu('add')}
      />
      <BottomNavigation.Item
        active={activeMenu === 'setting'}
        label="Setting"
        icon="settings"
        onPress={() => setActiveMenu('setting')}
      />
    </BottomNavigation>
  );
};
