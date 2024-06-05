import * as React from 'react';
import { BottomDrawer } from '@repo/design-system';

import { useUploadMedia } from '@/domain/media/hooks/useUploadMedia';
import { FileInput } from '@/modules/components';

type Props = {
  onUploadSuccess: (p: { mediumIDs: string[] }) => void;
};

export const MediaUploadDrawer = ({ onUploadSuccess }: Props) => {
  const inputRef = React.useRef<HTMLInputElement>(null);
  const [showUploadingDrawer, setShowUploadingDrawer] = React.useState(false);

  const { upload, statusLabel } = useUploadMedia();

  const handleChangeInput = (files: FileList | null) => {
    if (!files) {
      return;
    }

    setShowUploadingDrawer(true);

    upload({
      files,
      onSuccess: ({ mediumIDs }) => {
        onUploadSuccess({ mediumIDs });
        setShowUploadingDrawer(false);
      },
    });
  };

  return (
    <>
      <FileInput inputRef={inputRef} onChange={({ files }) => handleChangeInput(files)} />

      <BottomDrawer show={showUploadingDrawer} onClose={() => setShowUploadingDrawer(false)}>
        <BottomDrawer.Header onClose={() => setShowUploadingDrawer(false)}>
          Uploading media
        </BottomDrawer.Header>
        <BottomDrawer.Body>{statusLabel}</BottomDrawer.Body>
      </BottomDrawer>
    </>
  );
};
