import * as React from 'react';
import { useButton } from 'react-aria';
import { css } from '../../../styled-system/css';

import * as styles from '../../styles';

import { Icon, IconName } from '../Icon';

type Props = {
  label: React.ReactNode;
  icon: IconName;
  active?: boolean;
  onPress?: () => void;
};

export const Item = ({ label, icon, active = false, onPress }: Props) => {
  const ref = React.useRef<HTMLButtonElement>(null);
  const { buttonProps } = useButton({ onPress }, ref);

  const contentColor = active ? 'yellow.600' : 'gray.400';

  return (
    <button {...buttonProps} className={styles.classnames(styles.reset.button, Styles.item)}>
      <Icon size="S" color={contentColor} name={icon} />
      <span className={Styles.label} data-active={active}>
        {label}
      </span>
    </button>
  );
};

const Styles = {
  item: css({
    width: '3rem',
    height: '3rem',
    display: 'flex',
    flexDir: 'column',
    justifyContent: 'center',
    alignItems: 'center',
    gap: '0.125rem',
  }),
  label: css({
    fontSize: '0.625rem',
    color: 'gray.400',
    '&[data-active=true]': {
      color: 'yellow.600',
    },
  }),
};
