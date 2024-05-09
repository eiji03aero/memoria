import { css } from '../../../styled-system/css';

import * as styles from '../../styles';
import { Icon, IconName } from '../Icon';

type Props = {
  iconName: IconName;
  label: React.ReactNode;
  href: string;
  className?: string;
};

export const IconLinkCard = ({ iconName, label, href, className }: Props) => {
  return (
    <a href={href} className={styles.classnames(Styles.card, className)}>
      <Icon
        UNSAFE_className={Styles.icon}
        name={iconName}
        size="M"
        color="gray.900"
      />
      <span className={Styles.label}>{label}</span>
    </a>
  );
};

const Styles = {
  card: css({
    background: 'white',
    boxSizing: 'border-box',
    padding: '0.75rem',
    display: 'flex',
    flexDirection: 'column',
    gap: '0.25rem',
    borderRadius: '0.75rem',
    boxShadow: `0 0.25rem 1rem rgba(0,0,0,0.16)`,
  }),
  icon: css({
    flexShrink: 0,
  }),
  label: css({
    fontSize: '0.75rem',
    color: 'gray.900',
  }),
};
