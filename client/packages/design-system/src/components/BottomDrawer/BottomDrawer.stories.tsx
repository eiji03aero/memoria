import * as React from 'react';
import type { Meta, StoryFn } from '@storybook/react';

import { BottomDrawer } from './BottomDrawer';

const meta: Meta<typeof BottomDrawer> = {
  component: BottomDrawer,
  tags: ['autodocs'],
};

export default meta;

export const Default: StoryFn<typeof BottomDrawer> = () => {
  const [show, setShow] = React.useState(false);

  return (
    <>
      <button type="button" onClick={() => setShow(true)}>
        Show drawer
      </button>

      <BottomDrawer show={show} onClose={() => setShow(false)}>
        <BottomDrawer.Content>
          <BottomDrawer.Header onClose={() => setShow(false)}>Create album</BottomDrawer.Header>
          <BottomDrawer.Body>Body</BottomDrawer.Body>
          <BottomDrawer.Footer>Submit buttons will be here</BottomDrawer.Footer>
        </BottomDrawer.Content>
      </BottomDrawer>
    </>
  );
};

export const LotsOfContent: StoryFn<typeof BottomDrawer> = () => {
  const [show, setShow] = React.useState(false);

  return (
    <>
      <button type="button" onClick={() => setShow(true)}>
        Show drawer
      </button>

      <BottomDrawer show={show} onClose={() => setShow(false)}>
        <BottomDrawer.Content>
          <BottomDrawer.Header onClose={() => setShow(false)}>Create album</BottomDrawer.Header>
          <BottomDrawer.Body>
            {new Array(100).fill(0).map((_, idx) => (
              <p key={idx} style={{ marginBottom: '1rem' }}>
                Text aqui {idx}
              </p>
            ))}
          </BottomDrawer.Body>
          <BottomDrawer.Footer>Submit buttons will be here</BottomDrawer.Footer>
        </BottomDrawer.Content>
      </BottomDrawer>
    </>
  );
};

export const WithItem: StoryFn<typeof BottomDrawer> = () => {
  const [show, setShow] = React.useState(false);

  return (
    <>
      <button type="button" onClick={() => setShow(true)}>
        Show drawer
      </button>

      <BottomDrawer show={show} onClose={() => setShow(false)}>
        <BottomDrawer.Item onPress={console.log}>First</BottomDrawer.Item>
        <BottomDrawer.Item onPress={console.log}>Second</BottomDrawer.Item>
        <BottomDrawer.Item onPress={console.log}>
          I will rise and take back myself again
        </BottomDrawer.Item>
      </BottomDrawer>
    </>
  );
};
