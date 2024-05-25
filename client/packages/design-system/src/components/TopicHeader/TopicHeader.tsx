import { JSXElementConstructor } from 'react';
import { css } from '../../../styled-system/css';

import * as styles from '../../styles';
import { IconButton } from '../IconButton';

type Props = {
  as?: string | JSXElementConstructor<any>;
  label: React.ReactNode;
  trailing?: React.ReactNode;
  stickyTop?: boolean;
  showLoadingBar?: boolean;
  onBack?: () => void;
};

export const TopicHeader = ({
  as: Component = 'h1',
  label,
  trailing,
  stickyTop = false,
  showLoadingBar = false,
  onBack,
}: Props) => {
  return (
    <Component className={styles.classnames(Styles.header, stickyTop && Styles.stickyTop)}>
      {onBack && (
        <IconButton className={Styles.backIconButton} iconName="chevron-left" onPress={onBack} />
      )}
      <span className={Styles.label}>{label}</span>
      {!!trailing && <span className={Styles.trailing}>{trailing}</span>}

      <div
        className={styles.classnames(
          Styles.loadingBarContainer,
          showLoadingBar && Styles.loadingBarContainerShow,
        )}
      >
        <div className={Styles.loadingBarArea}>
          <div className={Styles.loadingBar} />
        </div>
      </div>
    </Component>
  );
};

const Styles = {
  header: css({
    display: 'flex',
    alignItems: 'center',
    position: 'relative',
    width: '100%',
    height: '3rem',
    padding: '0 0.5rem',
    bg: 'white',
    borderBottom: '1px solid',
    borderBottomColor: 'gray.200',
  }),
  stickyTop: css({
    position: 'sticky',
    top: 0,
    zIndex: 1,
  }),
  backIconButton: css({
    marginRight: '0.5rem',
  }),
  label: css({
    fontSize: '1.25rem',
    fontWeight: 'bold',
  }),
  trailing: css({
    marginLeft: 'auto',
  }),
  loadingBarContainer: css({
    position: 'absolute',
    top: '100%',
    left: '0',
    width: '100%',
    opacity: 0,
    pointerEvents: 'none',
    transition: 'opacity 250ms',
  }),
  loadingBarContainerShow: css({
    opacity: 1,
    pointerEvents: 'auto',
  }),
  loadingBarArea: css({
    position: 'relative',
    width: '100%',
  }),
  loadingBar: css({
    position: 'absolute',
    top: 0,
    height: '0.25rem',
    bg: 'yellow.400',
    animation: '2.5s linear 0s loadingBar infinite',
  }),
};
