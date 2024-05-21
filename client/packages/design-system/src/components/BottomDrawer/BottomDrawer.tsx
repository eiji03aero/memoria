import * as React from 'react';
import * as ReactDOM from 'react-dom';
import { css } from '../../../styled-system/css';

import * as styles from '../../styles';
import { Link } from '../Link';
import { Icon, IconName } from '../Icon';

type Props = {
  show: boolean;
  children: React.ReactNode;
  onClose: () => void;
};

export const BottomDrawer = ({ show, children, onClose }: Props) => {
  const content = (
    <>
      <div
        className={Styles.backdrop}
        onClick={d => {
          console.log('called from backdrop', d);
          onClose();
        }}
      />
      <div className={Styles.drawer}>{children}</div>
    </>
  );

  if (!show) {
    return null;
  }

  return ReactDOM.createPortal(content, document.body.querySelector('.JuTe6q_spectrum')!);
};

BottomDrawer.Content = ({ children }: { children: React.ReactNode }) => {
  return <div className={Styles.content}>{children}</div>;
};

BottomDrawer.Header = ({
  children,
  onClose,
}: {
  children: React.ReactNode;
  onClose?: () => void;
}) => {
  return (
    <div className={Styles.header}>
      {children}
      {onClose && <Link className={Styles.headerCancel} label="Cancel" onPress={onClose} />}
    </div>
  );
};

BottomDrawer.Body = ({
  className,
  children,
}: {
  className?: string;
  children: React.ReactNode;
}) => {
  return <div className={styles.classnames(Styles.body, className)}>{children}</div>;
};

BottomDrawer.Footer = ({ children }: { children?: React.ReactNode }) => {
  return <div className={Styles.footer}>{children}</div>;
};

BottomDrawer.Item = ({
  className,
  iconName,
  onPress,
  children,
}: {
  className?: string;
  iconName?: IconName;
  onPress?: () => void;
  children: React.ReactNode;
}) => {
  return (
    <div role="button" className={styles.classnames(Styles.item, className)} onClick={onPress}>
      {iconName && (
        <span className={Styles.itemIcon}>
          <Icon name={iconName} color="gray.600" size="S" />
        </span>
      )}
      {children}
    </div>
  );
};

const Styles = {
  backdrop: css({
    position: 'fixed',
    top: 0,
    left: 0,
    w: '100%',
    h: '100%',
    bg: 'black',
    opacity: 0.16,
  }),
  drawer: css({
    position: 'fixed',
    left: 0,
    bottom: 0,
    width: '100%',
    bg: 'white',
    borderTopLeftRadius: 'md',
    borderTopRightRadius: 'md',
  }),
  content: css({
    display: 'flex',
    flexDir: 'column',
    maxHeight: 'calc(100dvh - 4rem)',
  }),
  header: css({
    flexShrink: 0,
    display: 'flex',
    alignItems: 'center',
    height: '3rem',
    padding: '0 0.5rem',
    borderBottom: '1px solid',
    borderBottomColor: 'gray.200',
  }),
  headerCancel: css({
    marginLeft: 'auto',
  }),
  body: css({
    minHeight: 0,
    overflow: 'auto',
    padding: '0 0.5rem',
  }),
  footer: css({
    flexShrink: 0,
    display: 'flex',
    alignItems: 'center',
    height: '3.5rem',
    padding: '0 0.5rem',
    borderTop: '1px solid',
    borderTopColor: 'gray.200',
  }),
  item: css({
    display: 'flex',
    alignItems: 'center',
    width: '100%',
    height: '3rem',
    px: '0.5rem',
    fontSize: '1rem',
    color: 'black',
  }),
  itemIcon: css({
    display: 'flex',
    mr: '0.5rem',
  }),
};
