import * as React from 'react';

import * as config from '@/config';
import { axios, bareAxios, AxiosResponse } from '@/modules/lib/axios';
import { useMultipleRequestsStatus } from '@/modules/hooks/useMultipleRequestsStatus';

type RequestUploadURLsRequest = {
  file_names: string[];
  album_ids: string[];
};

type RequestUploadURLsResponse = {
  upload_urls: Array<{
    url: string;
    medium_id: string;
  }>;
};

const requestUploadURLs = (data: RequestUploadURLsRequest) =>
  axios.post<RequestUploadURLsRequest, AxiosResponse<RequestUploadURLsResponse>>(
    `${config.ApiHost}/api/auth/media/request-upload-urls`,
    data,
  );

const requestS3PutObject = (data: { url: string; file: File }) =>
  bareAxios.put(data.url, data.file, {
    headers: {
      'Content-Type': data.file.type,
    },
  });

const requestConfirmUploads = (data: { mediumIDs: string[] }) =>
  axios.post(`${config.ApiHost}/api/auth/media/confirm-uploads`, {
    medium_ids: data.mediumIDs,
  });

export const useUploadMedia = () => {
  const rs = useMultipleRequestsStatus();

  const upload = React.useCallback(
    async (params: { files: FileList; albumIDs: string[]; onSuccess: () => void }) => {
      const files = Array.from(params.files);

      // request creating media objects and signed urls
      rs.reset();
      rs.setStatus('preparing');
      const {
        data: { upload_urls },
      } = await requestUploadURLs({
        file_names: files.map(file => file.name),
        album_ids: params.albumIDs,
      });

      rs.setStatus('requesting');
      // with the signed urls, proceed with upload one by one
      rs.setTotalRequests(files.length);
      for (var i = 0; i < upload_urls.length; i++) {
        await requestS3PutObject({
          url: upload_urls[i]?.url!,
          file: files[i]!,
        });
        rs.incrementRequested();
        rs.setTotalRequested(prev => prev + 1);
      }

      rs.setStatus('completing');
      await requestConfirmUploads({ mediumIDs: upload_urls.map(uu => uu.medium_id) });

      rs.setStatus('completed');
      params.onSuccess?.();
    },
    [rs],
  );

  const statusLabel = React.useMemo(() => {
    if (rs.status === 'standby') {
      return null;
    }

    if (rs.status === 'preparing') {
      return 'Preparing to upload files ...';
    }

    if (rs.status === 'requesting') {
      return `Uploading files: ${rs.totalRequested} / ${rs.totalRequests}`;
    }

    if (rs.status === 'completing') {
      return `Completing the media upload ...`;
    }

    if (rs.status === 'completed') {
      return `Successfully completed media upload!`;
    }

    return null;
  }, [rs]);

  return {
    upload,
    statusLabel,
  };
};
