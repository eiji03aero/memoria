import { css } from '../../../styled-system/css';

import { Loader } from '../Loader';

export const LoadingBlock = () => {
  return (
    <div className={Styles.block}>
      <Loader />
    </div>
  );
};

const Styles = {
  block: css({
    display: 'flex',
    justifyContent: 'center',
    paddingY: '2rem',
  }),
};
