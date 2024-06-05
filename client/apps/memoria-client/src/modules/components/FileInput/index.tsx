import * as React from 'react';
import { Button } from '@repo/design-system';
import { css, cx } from '@/../styled-system/css';

type Props = {
  inputRef?: React.RefObject<HTMLInputElement>;
  hidden?: boolean;
  fileType?: 'media';
  onChange: (p: { files: FileList | null }) => void;
};

export const FileInput = ({
  inputRef: inputRefProp,
  hidden = false,
  fileType = 'media',
  onChange,
}: Props) => {
  const inputRef = React.useRef<HTMLInputElement>(null);

  const accept = React.useMemo(() => {
    if (fileType === 'media') {
      return 'image/*, image/heic, video/*';
    }

    return undefined;
  }, [fileType]);

  return (
    <>
      {!hidden && (
        <div>
          <Button
            variant="primary"
            onPress={() => {
              if (!inputRef.current) {
                return;
              }

              inputRef.current.value = '';
              inputRef.current.click();
            }}
          >
            Choose files
          </Button>
          {!!inputRef.current?.value && (
            <span className={css({ ml: '0.5rem' })}>{inputRef.current.value}</span>
          )}
        </div>
      )}
      <input
        ref={el => {
          if (!el) {
            return;
          }

          (inputRef as React.MutableRefObject<HTMLInputElement>).current = el;
          (inputRefProp as React.MutableRefObject<HTMLInputElement>).current = el;
        }}
        className={cx(Styles.hidden)}
        type="file"
        accept={accept}
        multiple
        onChange={e => onChange({ files: e.target.files })}
      />
    </>
  );
};

const Styles = {
  hidden: css({
    visibility: 'hidden',
    position: 'absolute',
  }),
};
