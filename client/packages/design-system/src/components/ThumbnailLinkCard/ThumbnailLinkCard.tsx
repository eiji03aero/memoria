import * as React from 'react';
import { useButton } from 'react-aria';
import { css } from '../../../styled-system/css';

import * as styles from '../../styles';
import { Icon } from '../Icon';
import { IconButton } from '../IconButton';

type Props = {
  src?: string;
  label: React.ReactNode;
  onPress: () => void;
  onOpenMenu?: () => void;
};

export const ThumbnailLinkCard = ({ label, onPress, onOpenMenu }: Props) => {
  return (
    <div
      role="button"
      className={styles.classnames(styles.reset.button, Styles.card)}
      onClick={onPress}
    >
      <span className={Styles.thumbnailArea}>
        <Icon name="image-album" color="gray.400" size="L" />
      </span>

      <span className={Styles.label}>{label}</span>

      {onOpenMenu && (
        <IconButton
          className={Styles.menuButton}
          iconName="more-vertical"
          onPress={onOpenMenu}
          iconSize="XS"
        />
      )}
    </div>
  );
};

const Styles = {
  card: css({
    display: 'flex',
    width: '100%',
    height: '4rem',
    padding: '0 0.5rem',
  }),
  thumbnailArea: css({
    display: 'inline-flex',
    justifyContent: 'center',
    alignItems: 'center',
    alignSelf: 'center',
    width: '4rem',
    height: '3rem',
    bg: 'gray.200',
    border: '1px solid',
    borderColor: 'gray.400',
    borderRadius: 'md',
    marginRight: '1rem',
  }),
  label: css({
    fontSize: '1rem',
    color: 'black',
    alignSelf: 'center',
  }),
  menuButton: css({
    ml: 'auto',
    alignSelf: 'center',
  }),
};
