import * as React from 'react';
import { useTranslation } from 'react-i18next';

import * as config from '@/config';
import { axios } from '@/modules/lib/axios';
import { useMultipleRequestsStatus } from '@/modules/hooks/useMultipleRequestsStatus';

const request = (id: string) => axios.delete(`${config.ApiHost}/api/auth/media/${id}`);

export const useDeleteMedia = () => {
  const { t } = useTranslation();
  const rs = useMultipleRequestsStatus();

  const deleteMedia = React.useCallback(
    async (params: { ids: string[]; onSuccess: () => void }) => {
      rs.setStatus('requesting');
      rs.reset();
      rs.setTotalRequests(params.ids.length);

      for (let i = 0; i < params.ids.length; i++) {
        await request(params.ids[i]!);
        rs.incrementRequested();
      }

      rs.setStatus('completed');
      params.onSuccess();
    },
    [rs],
  );

  const statusLabel = React.useMemo(() => {
    if (rs.status === 'requesting') {
      return `${t('w.deleting-data', { data: t('w.media').toLowerCase() })}: ${rs.totalRequested} / ${rs.totalRequests}`;
    }

    if (rs.status === 'completed') {
      return t('w.deletion-completed');
    }

    return null;
  }, [rs]);

  return {
    deleteMedia,
    statusLabel,
  };
};
