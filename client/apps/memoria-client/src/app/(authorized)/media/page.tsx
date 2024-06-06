'use client';

import { useTranslation } from 'react-i18next';

import { MediaScreen } from '@/domain/media/standalone/MediaScreen';

export default function AllMedia() {
  const { t } = useTranslation();
  return <MediaScreen title={t('w.all-data', { data: t('w.media').toLowerCase() })} />;
}
