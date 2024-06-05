import * as React from 'react';
import { parseISO, format } from 'date-fns';

export const useDateFormat = () => {
  const pf = React.useMemo(
    () => ({
      fullDateDOW: (str: string) => {
        return format(parseISO(str), 'yyyy-MM-dd HH:mm:ss (eee)');
      },
    }),
    [],
  );

  return React.useMemo(() => ({ pf }), [pf]);
};
