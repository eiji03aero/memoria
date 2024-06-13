import { datadogLogs } from '@datadog/browser-logs';

import * as config from '@/config';

export const init = () => {
  if (!config.DatadogClientToken) {
    return;
  }

  datadogLogs.init({
    clientToken: config.DatadogClientToken,
    site: 'ap1.datadoghq.com',
    forwardErrorsToLogs: true,
    sessionSampleRate: 100,
  });
};
