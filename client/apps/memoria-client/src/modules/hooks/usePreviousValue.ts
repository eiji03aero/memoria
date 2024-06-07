import * as React from 'react';

export const usePreviousValue = <T>(
  v: T,
): {
  previousValue: T;
  changed: boolean;
} => {
  const ref = React.useRef<T>(v);

  React.useEffect(() => {
    ref.current = v;
  }, [v]);

  return {
    previousValue: ref.current,
    changed: v === ref.current,
  };
};
