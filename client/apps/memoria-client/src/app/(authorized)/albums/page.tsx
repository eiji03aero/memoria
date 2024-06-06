'use client';

import * as React from 'react';
import { useRouter } from 'next/navigation';
import { useTranslation } from 'react-i18next';
import {
  TightLayoutCard,
  ThumbnailLinkCard,
  IconButton,
  BottomDrawer,
  Button,
} from '@repo/design-system';

import { TopicHeader } from '@/domain/common/standalone/TopicHeader';
import { CreateAlbumDrawer } from '@/domain/album/standalone/CreateAlbumDrawer';
import { Album } from '@/domain/common/interfaces/api';
import { useAlbums } from '@/domain/album/hooks/useAlbums';
import { useDeleteAlbum } from '@/domain/album/hooks/useDeleteAlbum';

export default function Albums() {
  const { t } = useTranslation();
  const [showCreateAlbumDrawer, setShowCreateAlbumDrawer] = React.useState(false);
  const [showDeleteAlbumDrawer, setShowDeleteAlbumDrawer] = React.useState(false);
  const [showAlbumMenuDrawer, setShowAlbumMenuDrawer] = React.useState(false);
  const [selectedAlbum, setSelectedAlbum] = React.useState<Album | null>(null);

  const { albums } = useAlbums();
  const { deleteAlbum } = useDeleteAlbum();
  const router = useRouter();

  const handleOpenMenu = React.useCallback((album: Album) => {
    setSelectedAlbum(album);
    setShowAlbumMenuDrawer(true);
  }, []);

  const handleDelete = React.useCallback(() => {
    if (!selectedAlbum) {
      return;
    }

    deleteAlbum(selectedAlbum.id, {
      onSuccess: () => {
        setShowDeleteAlbumDrawer(false);
      },
    });
  }, [selectedAlbum, deleteAlbum, setShowDeleteAlbumDrawer]);

  return (
    <>
      <TopicHeader
        label={t('w.albums')}
        trailing={
          <IconButton
            variant="elevated"
            iconName="add"
            onPress={() => {
              setShowCreateAlbumDrawer(true);
            }}
          />
        }
      />

      <TightLayoutCard.Background>
        <TightLayoutCard>
          <ThumbnailLinkCard label={t('w.all')} onPress={() => router.push(`/media`)} />
        </TightLayoutCard>

        <TightLayoutCard>
          {albums?.map(album => (
            <ThumbnailLinkCard
              key={album.id}
              label={album.name}
              onPress={() => router.push(`/albums/${album.id}`)}
              onOpenMenu={() => handleOpenMenu(album)}
            />
          ))}
        </TightLayoutCard>
      </TightLayoutCard.Background>

      <CreateAlbumDrawer
        show={showCreateAlbumDrawer}
        onClose={() => setShowCreateAlbumDrawer(false)}
      />

      <BottomDrawer show={showAlbumMenuDrawer} onClose={() => setShowAlbumMenuDrawer(false)}>
        <BottomDrawer.Header onClose={() => setShowDeleteAlbumDrawer(false)}>
          {t('w.data-menu', { data: t('w.album') })}
        </BottomDrawer.Header>
        <BottomDrawer.Item
          iconName="delete"
          onPress={() => {
            setShowAlbumMenuDrawer(false);
            setShowDeleteAlbumDrawer(true);
          }}
        >
          {t('w.delete')}
        </BottomDrawer.Item>
        <BottomDrawer.Footer />
      </BottomDrawer>

      <BottomDrawer show={showDeleteAlbumDrawer} onClose={() => setShowDeleteAlbumDrawer(false)}>
        <BottomDrawer.Header onClose={() => setShowDeleteAlbumDrawer(false)}>
          {t('w.delete-data', { data: t('w.album').toLowerCase() })}
        </BottomDrawer.Header>
        <BottomDrawer.Body>
          {t('s.are-you-sure-to-delete-this-data', { data: t('w.album').toLowerCase() })}
          <br />
          {selectedAlbum?.name}
        </BottomDrawer.Body>
        <BottomDrawer.Footer>
          <Button variant="primary" onPress={() => handleDelete()}>
            {t('w.delete')}
          </Button>
        </BottomDrawer.Footer>
      </BottomDrawer>
    </>
  );
}
