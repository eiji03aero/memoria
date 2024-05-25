import * as React from 'react';

type Status = 'standby' | 'preparing' | 'requesting' | 'completing' | 'completed';

export const useMultipleRequestsStatus = () => {
  const [status, setStatus] = React.useState<Status>('standby');
  const [totalRequests, setTotalRequests] = React.useState(0);
  const [totalRequested, setTotalRequested] = React.useState(0);

  const incrementRequested = React.useCallback(() => setTotalRequested(prev => prev + 1), []);
  const reset = React.useCallback(() => {
    setTotalRequested(0);
    setTotalRequests(0);
  }, []);

  return React.useMemo(
    () => ({
      status,
      setStatus,
      totalRequests,
      setTotalRequests,
      totalRequested,
      setTotalRequested,
      incrementRequested,
      reset,
    }),
    [
      status,
      setStatus,
      totalRequests,
      setTotalRequests,
      totalRequested,
      setTotalRequested,
      incrementRequested,
      reset,
    ],
  );
};
