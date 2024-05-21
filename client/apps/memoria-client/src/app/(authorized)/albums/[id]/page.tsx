'use client';

import { useAlbum } from '@/domain/album/hooks/useAlbum';
import { MediaScreen } from '@/app/(authorized)/_standalone/MediaScreen';

export default function Album({ params }: { params: { id: string } }) {
  const { album } = useAlbum({ id: params.id });

  return <MediaScreen title={album?.name ?? 'Loading'} albumID={params.id} />;
}
