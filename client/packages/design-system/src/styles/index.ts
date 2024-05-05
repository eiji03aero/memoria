import classnames from 'classnames';
import { ColorValue } from '@react-types/shared';

import { css } from '../../styled-system/css';

export { useMediaQueries } from './useMediaQueries';

export { classnames, css };

export const BreakPoints = {
  Mobile: 768,
};

export const MediaQueries = {
  Mobile: `(min-width: ${BreakPoints.Mobile})`,
};

export const reset = {
  button: css({
    bg: 'transparent',
    border: 'none',
    cursor: 'pointer',
    outline: 'none',
    appearance: 'none',
  }),
} as const;

const SizeKeys = [
  'size-0',
  'size-10',
  'size-25',
  'size-40',
  'size-50',
  'size-65',
  'size-75',
  'size-85',
  'size-100',
  'size-115',
  'size-125',
  'size-130',
  'size-150',
  'size-160',
  'size-175',
  'size-200',
  'size-225',
  'size-250',
  'size-275',
  'size-300',
  'size-325',
  'size-350',
  'size-400',
  'size-450',
  'size-500',
  'size-550',
  'size-600',
  'size-675',
  'size-700',
  'size-800',
  'size-900',
  'size-1000',
  'size-1200',
  'size-1250',
  'size-1600',
  'size-1700',
  'size-2000',
  'size-2400',
  'size-3000',
  'size-3400',
  'size-3600',
  'size-4600',
  'size-5000',
  'size-6000',
] as const;

type SizeKey = (typeof SizeKeys)[number];

export const CSSVar = {
  size: (size: SizeKey) =>
    `var(--spectrum-global-dimension-${size}, var(--spectrum-alias-${size}))`,
  color: (color: ColorValue) =>
    `var(--spectrum-legacy-color-${color}, var(--spectrum-global-color-${color}))`,
};
