'use client';

import { useSearchParams, useRouter } from 'next/navigation';

import { MediaGalleryScreen } from '@/domain/media/standalone/MediaGalleryScreen';

export default function MediaGallery() {
  const searchParams = useSearchParams();
  const router = useRouter();
  const albumID = searchParams.get('album_id') ?? undefined;
  const initialMediumID = searchParams.get('initial_medium_id');
  if (!initialMediumID) {
    throw new Error('initial medium id is not given');
  }

  return (
    <MediaGalleryScreen
      albumID={albumID}
      onBack={() => router.back()}
      initialMediumID={initialMediumID}
    />
  );
}
