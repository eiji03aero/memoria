import type { Meta, StoryFn } from '@storybook/react';
import { css } from '../../../styled-system/css';

import { Icon, IconNames } from './Icon';

const meta: Meta<typeof Icon> = {
  component: Icon,
  tags: ['autodocs'],
};

export default meta;

export const List: StoryFn<typeof Icon> = () => {
  return (
    <div
      className={css({
        display: 'flex',
        flexWrap: 'wrap',
        gap: '0.5rem',
      })}
    >
      {IconNames.map(name => (
        <div
          className={css({
            display: 'flex',
            flexDir: 'column',
            alignContent: 'center',
            justifyContent: 'center',
            gap: '1rem',
            border: '1px solid gray.200',
            borderRadius: 'md',
            width: '2xl',
            height: '2xl',
          })}
        >
          <Icon name={name} color="gray.900" size="XL" />
          <span
            className={css({
              fontSize: '1rem',
              color: 'gray.900',
            })}
          >
            {name}
          </span>
        </div>
      ))}
    </div>
  );
};
