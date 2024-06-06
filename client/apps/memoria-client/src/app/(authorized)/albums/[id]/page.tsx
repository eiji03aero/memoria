'use client';

import { useTranslation } from 'react-i18next';

import { useAlbum } from '@/domain/album/hooks/useAlbum';
import { MediaScreen } from '@/domain/media/standalone/MediaScreen';

export default function Album({ params }: { params: { id: string } }) {
  const { t } = useTranslation();
  const { album } = useAlbum({ id: params.id });

  return <MediaScreen title={album?.name ?? t('w.loading')} albumID={params.id} />;
}
