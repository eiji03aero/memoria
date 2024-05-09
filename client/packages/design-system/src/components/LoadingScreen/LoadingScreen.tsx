import { css } from '../../../styled-system/css';

import { Loader } from '../Loader';

export const LoadingScreen = () => {
  return (
    <div className={Styles.screen}>
      <Loader />
    </div>
  );
};

const Styles = {
  screen: css({
    display: 'flex',
    justifyContent: 'center',
    paddingY: '2rem',
  }),
};
