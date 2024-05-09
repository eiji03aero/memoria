import classnames from 'classnames';

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
