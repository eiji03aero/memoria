import { css } from '../../../styled-system/css';

import * as styles from '../../styles';
import { Icon, IconName, IconSize } from '../Icon';

type Props = {
  variant?: 'default' | 'elevated';
  iconName: IconName;
  iconSize?: IconSize;
  className?: string;
  onPress: () => void;
};

export const IconButton = ({
  variant = 'default',
  iconName,
  iconSize = 'S',
  className,
  onPress,
}: Props) => {
  return (
    <button
      type="button"
      className={styles.classnames(
        styles.reset.button,
        Styles.button,
        Styles.variant[variant].bg,
        className,
      )}
      onClick={e => {
        e.stopPropagation();
        onPress();
      }}
    >
      <Icon name={iconName} color={Styles.variant[variant].icon} size={iconSize} />
    </button>
  );
};

const Styles = {
  button: css({
    display: 'inline-flex',
    alignItems: 'center',
    justifyContent: 'center',
    w: '2rem',
    h: '2rem',
    borderRadius: '50%',
    '& + &': {
      ml: '0.5rem',
    },
  }),
  variant: {
    default: {
      icon: 'gray.600' as const,
      bg: css({
        bg: 'transparent',
      }),
    },
    elevated: {
      icon: 'gray.600' as const,
      bg: css({
        bg: 'gray.200',
      }),
    },
  },
};
