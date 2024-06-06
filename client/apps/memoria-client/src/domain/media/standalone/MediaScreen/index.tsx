'use client';

import * as React from 'react';
import { useTranslation } from 'react-i18next';
import { useRouter } from 'next/navigation';
import { IconButton, MediaGrid, BottomDrawer, Button } from '@repo/design-system';
import { css } from '@/../styled-system/css';

import { TopicHeader } from '@/domain/common/standalone/TopicHeader';
import { Medium } from '@/domain/common/interfaces/api';
import { MediaGridTile } from '@/domain/media/components/MediaGridTile';
import { AlbumSelect } from '@/domain/album/standalone/AlbumSelect';
import { useUploadMedia } from '@/domain/media/hooks/useUploadMedia';
import { useMedia } from '@/domain/media/hooks/useMedia';
import { useDeleteMedia } from '@/domain/media/hooks/useDeleteMedia';
import { useAddMediaToAlbums } from '@/domain/album/hooks/useAddMediaToAlbums';
import { useRemoveMediaFromAlbums } from '@/domain/album/hooks/useRemoveMediaFromAlbums';
import { FileInput } from '@/modules/components';
import * as utils from '@/modules/utils';

type Props = {
  title: string;
  albumID?: string;
};

export const MediaScreen = ({ title, albumID }: Props) => {
  const { t } = useTranslation();
  const fileInputRef = React.useRef<HTMLInputElement>(null);
  const router = useRouter();
  const { upload, statusLabel } = useUploadMedia();
  const [showUploadingDrawer, setShowUploadingDrawer] = React.useState(false);
  const { media, refetch } = useMedia({ albumID });
  const { deleteMedia, statusLabel: deleteStatusLabel } = useDeleteMedia();
  const [showDeletingDrawer, setShowDeletingDrawer] = React.useState(false);
  const { addMediaToAlbums } = useAddMediaToAlbums();
  const [showAddMediaToAlbumsDrawer, setShowAddMediaToAlbumsDrawer] = React.useState(false);
  const [selectedAlbumIDs, setSelectedAlbumIDs] = React.useState<string[]>([]);
  const { removeMediaFromAlbums } = useRemoveMediaFromAlbums();
  const [showUploadConfirmDrawer, setShowUploadConfirmDrawer] = React.useState(false);

  const [mode, setMode] = React.useState<'default' | 'selecting'>('default');
  const [selectedIds, setSelectedIds] = React.useState<string[]>([]);

  const startSelectingFilesToUpload = () => {
    fileInputRef.current?.click();
  };

  const files = (fileInputRef.current as HTMLInputElement)?.files;

  const handleChangeFileInput = () => {
    if (files === null) {
      return;
    }

    setShowUploadConfirmDrawer(true);
  };

  const handleUploadMedia = () => {
    if (files === null) {
      return;
    }

    setSelectedIds([]);
    setSelectedAlbumIDs([]);
    upload({ files, albumIDs: selectedAlbumIDs, onSuccess: () => refetch() });
    setShowUploadConfirmDrawer(false);
    setShowUploadingDrawer(true);
  };

  const handlePressTile = React.useCallback(
    (medium: Medium) => {
      if (mode === 'default') {
        // TODO: navigate to medium detail
        const qs = utils.buildQueryParams({ album_id: albumID, initial_medium_id: medium.id });
        router.push(`/media/gallery?${qs}`);
      } else if (mode === 'selecting') {
        setSelectedIds(prev => {
          if (prev.includes(medium.id)) {
            return prev.filter(id => id !== medium.id);
          } else {
            return prev.concat(medium.id);
          }
        });
      }
    },
    [mode, router, albumID],
  );

  const handleDelete = React.useCallback(() => {
    if (selectedIds.length === 0) {
      return;
    }

    setShowDeletingDrawer(true);
    setMode('default');
    deleteMedia({
      ids: selectedIds,
      onSuccess: () => {
        refetch();
      },
    });
  }, [selectedIds, deleteMedia, refetch]);

  const handleAddMediaToAlbums = React.useCallback(() => {
    setMode('default');
    setSelectedIds([]);
    setSelectedAlbumIDs([]);
    addMediaToAlbums(
      {
        albumIDs: selectedAlbumIDs,
        mediumIDs: selectedIds,
      },
      {
        onSuccess: () => {
          setShowAddMediaToAlbumsDrawer(false);
        },
      },
    );
  }, [selectedIds, selectedAlbumIDs, addMediaToAlbums]);

  const handleRemoveMediaFromAlbums = React.useCallback(() => {
    if (!albumID) {
      return;
    }

    setMode('default');
    setSelectedIds([]);
    setSelectedAlbumIDs([]);
    removeMediaFromAlbums({
      albumIDs: [albumID],
      mediumIDs: selectedIds,
    });
  }, [albumID, selectedIds, removeMediaFromAlbums]);

  return (
    <>
      <FileInput hidden inputRef={fileInputRef} onChange={handleChangeFileInput} />

      <TopicHeader
        label={title}
        stickyTop
        trailing={
          <>
            {mode === 'default' && (
              <>
                <IconButton
                  variant="elevated"
                  iconName="add"
                  onPress={startSelectingFilesToUpload}
                />
                <IconButton
                  variant="elevated"
                  iconName="select-add"
                  onPress={() => setMode('selecting')}
                />
              </>
            )}
            {mode === 'selecting' && (
              <>
                <IconButton variant="elevated" iconName="delete" onPress={handleDelete} />
                <IconButton
                  variant="elevated"
                  iconName="folder-add"
                  onPress={() => setShowAddMediaToAlbumsDrawer(true)}
                />
                <IconButton
                  variant="elevated"
                  iconName="folder-remove"
                  onPress={handleRemoveMediaFromAlbums}
                />
                <IconButton
                  variant="elevated"
                  iconName="cancel"
                  onPress={() => {
                    setMode('default');
                    setSelectedIds([]);
                    setSelectedAlbumIDs([]);
                  }}
                />
              </>
            )}
          </>
        }
        onBack={() => router.push('/albums')}
      />

      <MediaGrid>
        {media?.map(medium => {
          return (
            <MediaGridTile
              key={medium.id}
              selected={selectedIds.includes(medium.id)}
              medium={medium}
              onPress={() => handlePressTile(medium)}
            />
          );
        })}
      </MediaGrid>

      <BottomDrawer
        show={showUploadConfirmDrawer}
        onClose={() => setShowUploadConfirmDrawer(false)}
      >
        <BottomDrawer.Header onClose={() => setShowUploadConfirmDrawer(false)}>
          {t('p.media-screen.before-upload-selected-media')}
        </BottomDrawer.Header>
        <BottomDrawer.Body>
          <AlbumSelect onChange={values => setSelectedAlbumIDs(values.map(v => v.value))} />
        </BottomDrawer.Body>
        <BottomDrawer.Footer>
          <Button variant="primary" onPress={handleUploadMedia}>
            {t('w.upload')}
          </Button>
        </BottomDrawer.Footer>
      </BottomDrawer>

      <BottomDrawer show={showUploadingDrawer} onClose={() => setShowUploadingDrawer(false)}>
        <BottomDrawer.Header onClose={() => setShowUploadingDrawer(false)}>
          {t('w.uploading-data', { data: t('w.media') })}
        </BottomDrawer.Header>
        <BottomDrawer.Body>
          <div className={Styles.drawerBody}>{statusLabel}</div>
        </BottomDrawer.Body>
      </BottomDrawer>

      <BottomDrawer show={showDeletingDrawer} onClose={() => setShowDeletingDrawer(false)}>
        <BottomDrawer.Header onClose={() => setShowDeletingDrawer(false)}>
          {t('w.deleting-data', { data: t('w.media').toLowerCase() })}
        </BottomDrawer.Header>
        <BottomDrawer.Body>
          <div className={Styles.drawerBody}>{deleteStatusLabel}</div>
        </BottomDrawer.Body>
      </BottomDrawer>

      <BottomDrawer
        show={showAddMediaToAlbumsDrawer}
        onClose={() => setShowAddMediaToAlbumsDrawer(false)}
      >
        <BottomDrawer.Header onClose={() => setShowAddMediaToAlbumsDrawer(false)}>
          {t('p.media-screen.adding-media-to-albums')}
        </BottomDrawer.Header>
        <BottomDrawer.Body>
          <AlbumSelect onChange={values => setSelectedAlbumIDs(values.map(v => v.value))} />
        </BottomDrawer.Body>
        <BottomDrawer.Footer>
          <Button variant="primary" onPress={handleAddMediaToAlbums}>
            {t('w.save')}
          </Button>
        </BottomDrawer.Footer>
      </BottomDrawer>
    </>
  );
};

const Styles = {
  fileInput: css({
    visibility: 'hidden',
    position: 'absolute',
  }),
  drawerBody: css({
    py: '1rem',
  }),
};
