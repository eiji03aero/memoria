import * as React from 'react';

import '../src/index.css';
import { Provider } from '../src/Provider';

/** @type { import('@storybook/react').Preview } */
const preview = {
  parameters: {
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/i,
      },
    },
  },
  decorators: [
    Story => (
      <Provider>
        <Story />
      </Provider>
    ),
  ],
};

export default preview;
