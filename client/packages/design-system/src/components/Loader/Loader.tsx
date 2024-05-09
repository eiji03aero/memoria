import { css } from '../../../styled-system/css';

export const Loader = () => {
  return <div className={Styles.loader} />;
};

const Styles = {
  loader: css({
    width: '3rem',
    aspectRatio: 1,
    borderRadius: '50%',
    border: '0.25rem solid',
    borderColor: 'gray.100',
    borderRightColor: 'yellow.400',
    animation: 'spin 1s infinite linear',
  }),
};
