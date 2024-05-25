import { css } from '@/../styled-system/css';

type Props = {
  src: string;
};

export const VideoPanel = ({ src }: Props) => {
  return (
    <div className={Styles.panel}>
      <video className={Styles.video} src={src} />
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
  video: css({
    width: '100%',
    objectFit: 'contain',
  }),
};
