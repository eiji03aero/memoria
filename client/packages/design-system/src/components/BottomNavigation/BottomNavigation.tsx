import { css } from '../../../styled-system/css';

import { Item } from './Item';

export const BottomNavigation = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  return <nav className={Styles.navigation}>{children}</nav>;
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
};

BottomNavigation.Item = Item;
