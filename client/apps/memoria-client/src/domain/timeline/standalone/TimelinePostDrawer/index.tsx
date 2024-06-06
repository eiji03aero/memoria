import * as React from 'react';
import { useTranslation } from 'react-i18next';
import { Button, BottomDrawer, TextArea } from '@repo/design-system';
import { css } from '@/../styled-system/css';

import { useCreateTimelinePost } from '@/domain/timeline/hooks/useCreateTimelinePost';
import { MediaUploadDrawer } from '@/domain/media/standalone/MediaUploadDrawer';

type Props = {
  show: boolean;
  onClose: () => void;
};

export const TimelinePostDrawer = ({ show, onClose }: Props) => {
  const { t } = useTranslation();
  const [content, setContent] = React.useState('');
  const [uploadedMediumIDs, setUploadedMediumIDs] = React.useState<string[]>([]);

  const { createPost } = useCreateTimelinePost();

  const handleCreate = async () => {
    createPost({ content, mediumIDs: uploadedMediumIDs }, { onSuccess: onClose });
  };

  return (
    <BottomDrawer show={show} onClose={onClose}>
      <BottomDrawer.Header onClose={onClose}>{t('s.create-a-post')}</BottomDrawer.Header>
      <BottomDrawer.Body>
        <TextArea
          label={t('w.content')}
          value={content}
          height="240px"
          width="100%"
          onChange={v => setContent(v)}
        />

        <div className={css({ my: '0.5rem' })}>
          <MediaUploadDrawer onUploadSuccess={({ mediumIDs }) => setUploadedMediumIDs(mediumIDs)} />
        </div>
      </BottomDrawer.Body>
      <BottomDrawer.Footer>
        <Button variant="primary" onPress={() => handleCreate()}>
          {t('w.submit')}
        </Button>
      </BottomDrawer.Footer>
    </BottomDrawer>
  );
};
