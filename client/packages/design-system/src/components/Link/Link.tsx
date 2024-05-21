import * as React from 'react';
import { useButton } from 'react-aria';
import { css } from '../../../styled-system/css';

import * as styles from '../../styles';

type Props = {
  label?: React.ReactNode;
  className?: string;
  onPress?: () => void;
};

export const Link = ({ label, className, onPress }: Props) => {
  const ref = React.useRef<HTMLButtonElement>(null);
  const { buttonProps } = useButton({ onPress }, ref);
  return (
    <button
      type="button"
      {...buttonProps}
      className={styles.classnames(Styles.link, className)}
    >
      {label}
    </button>
  );
};

const Styles = {
  link: css({
    color: 'blue.400',
    fontSize: '1rem',
  }),
};
