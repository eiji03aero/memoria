import { css } from '../../../styled-system/css';
import * as styles from '../../styles';

type Props = {
  className?: string;
  children: React.ReactNode;
};

export const TightLayoutCard = ({ className, children }: Props) => {
  return <div className={styles.classnames(Styles.card, className)}>{children}</div>;
};

type BackgroundProps = {
  children: React.ReactNode;
};

TightLayoutCard.Background = ({ children }: BackgroundProps) => {
  return <div className={Styles.background}>{children}</div>;
};

const Styles = {
  card: css({
    width: '100%',
    background: 'white',
  }),
  background: css({
    display: 'flex',
    flexDir: 'column',
    gap: '0.125rem',
    minHeight: '100%',
    paddingY: '0.125rem',
    background: 'gray.300',
    boxSizing: 'border-box',
  }),
};
