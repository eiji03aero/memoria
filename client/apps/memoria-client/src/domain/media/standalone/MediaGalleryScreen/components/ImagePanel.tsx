import { css } from '@/../styled-system/css';

type Props = {
  src: string;
};

export const ImagePanel = ({ src }: Props) => {
  return (
    <div className={Styles.panel}>
      <img className={Styles.img} src={src} />
    </div>
  );
};

const Styles = {
  panel: css({
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    width: '100%',
    height: 'calc(100dvh - 3.125rem)',
  }),
  img: css({
    width: '100%',
    objectFit: 'contain',
  }),
};
