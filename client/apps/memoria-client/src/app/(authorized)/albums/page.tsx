'use client';

import * as React from 'react';
import { useRouter } from 'next/navigation';
import {
  TightLayoutCard,
  ThumbnailLinkCard,
  IconButton,
  BottomDrawer,
  Button,
} from '@repo/design-system';

import { TopicHeader } from '@/domain/common/standalone/TopicHeader';
import { Album } from '@/domain/common/interfaces/api';
import { useAlbums } from '@/domain/album/hooks/useAlbums';
import { useDeleteAlbum } from '@/domain/album/hooks/useDeleteAlbum';

import { CreateAlbumDrawer } from '@/app/(authorized)/albums/_standalone/CreateAlbumDrawer';

export default function Albums() {
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
        label="Albums"
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
          <ThumbnailLinkCard label="All" onPress={() => router.push(`/media`)} />
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
          Album menu
        </BottomDrawer.Header>
        <BottomDrawer.Item
          iconName="delete"
          onPress={() => {
            setShowAlbumMenuDrawer(false);
            setShowDeleteAlbumDrawer(true);
          }}
        >
          Delete
        </BottomDrawer.Item>
        <BottomDrawer.Footer />
      </BottomDrawer>

      <BottomDrawer show={showDeleteAlbumDrawer} onClose={() => setShowDeleteAlbumDrawer(false)}>
        <BottomDrawer.Header onClose={() => setShowDeleteAlbumDrawer(false)}>
          Delete album
        </BottomDrawer.Header>
        <BottomDrawer.Body>
          Are you sure to delete this album?
          <br />
          {selectedAlbum?.name}
        </BottomDrawer.Body>
        <BottomDrawer.Footer>
          <Button variant="primary" onPress={() => handleDelete()}>
            Delete
          </Button>
        </BottomDrawer.Footer>
      </BottomDrawer>
    </>
  );
}
