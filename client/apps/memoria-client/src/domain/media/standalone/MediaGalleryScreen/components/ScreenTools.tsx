import { css } from '@/../styled-system/css';
import { IconButton } from '@repo/design-system';

type Props = {
  show: boolean;
  onBack: () => void;
};

export const ScreenTools = ({ show, onBack }: Props) => {
  if (!show) {
    return;
  }

  return (
    <>
      <div className={Styles.top}>
        <div className={Styles.toolbar}>
          <IconButton iconName="chevron-left" onPress={onBack} />
        </div>
      </div>
    </>
  );
};

const Styles = {
  top: css({
    top: 0,
    left: 0,
    right: 0,
    position: 'absolute',
  }),
  toolbar: css({
    display: 'flex',
    alignItems: 'center',
    height: '3rem',
    padding: '0 0.5rem',
  }),
};
