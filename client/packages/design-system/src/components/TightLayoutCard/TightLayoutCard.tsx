import { css } from '../../../styled-system/css';
import * as styles from '../../styles';

type Props = {
  className?: string;
  children: React.ReactNode;
};

export const TightLayoutCard = ({ className, children }: Props) => {
  return (
    <div className={styles.classnames(Styles.card, className)}>{children}</div>
  );
};

type BackgroundProps = {
  children: React.ReactNode;
};

TightLayoutCard.Background = ({ children }: BackgroundProps) => {
  return <div className={Styles.background}>{children}</div>;
};

const Styles = {
  card: css({
    background: 'white',
    margin: '0.125rem 0',
  }),
  background: css({
    height: '100%',
    paddingY: '0.125rem',
    background: 'gray.300',
    boxSizing: 'border-box',
  }),
};
