import { css, cx } from '../../../styled-system/css';

import { Item } from './Item';

export const BottomNavigation = ({ children }: { children: React.ReactNode }) => {
  return (
    <nav className={cx(Styles.navigation, (window as any).navigator.standalone && Styles.pb)}>
      {children}
    </nav>
  );
};

const Styles = {
  navigation: css({
    width: '100%',
    display: 'flex',
    justifyContent: 'space-around',
    borderTop: '1px solid',
    borderColor: 'gray.200',
    backgroundColor: 'white',
  }),
  pb: css({
    pb: '1rem',
  }),
};

BottomNavigation.Item = Item;
