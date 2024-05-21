import * as React from 'react';
import { useButton } from 'react-aria';
import { css } from '../../../styled-system/css';

import * as styles from '../../styles';
import { Icon } from '../Icon';

type Props = {
  children: React.ReactNode;
};

export const MediaGrid = ({ children }: Props) => {
  return <div className={Styles.grid}>{children}</div>;
};

MediaGrid.Tile = ({
  src,
  selected,
  onImgError,
  onPress,
}: {
  src?: string;
  selected?: boolean;
  onImgError?: () => void;
  onPress?: () => void;
}) => {
  const ref = React.useRef<HTMLButtonElement>(null);
  const { buttonProps } = useButton({ onPress }, ref);
  return (
    <button
      ref={ref}
      type="button"
      className={styles.classnames(styles.reset.button, Styles.tile)}
      {...buttonProps}
    >
      <div className={Styles.tileContent}>
        {src && <img className={Styles.tileImage} onError={onImgError} src={src} />}
        {!src && <Icon size="XL" color="gray.400" name="image-album" />}
      </div>

      {selected && <div className={Styles.selected} />}
    </button>
  );
};

const Styles = {
  grid: css({
    display: 'grid',
    gridTemplateColumns: 'repeat(3, 1fr)',
    gap: '0.5rem',
    p: '0.5rem',
  }),
  tile: css({
    height: 0,
    pb: '100%',
    position: 'relative',
  }),
  selected: css({
    position: 'absolute',
    top: 0,
    left: 0,
    width: '100%',
    height: '100%',
    border: '0.25rem solid',
    borderColor: 'yellow.400',
  }),
  tileContent: css({
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    position: 'absolute',
    top: 0,
    left: 0,
    width: '100%',
    height: '100%',
    bg: 'gray.200',
  }),
  tileImage: css({
    width: '100%',
    height: '100%',
    objectFit: 'contain',
    objectPosition: 'center',
  }),
};
