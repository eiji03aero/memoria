import * as React from 'react';
import { TopicHeader as Base } from '@repo/design-system';
import { useIsRequesting } from '@/modules/hooks/useIsRequesting';

export const TopicHeader = (props: React.ComponentProps<typeof Base>) => {
  const isRequesting = useIsRequesting();
  return <Base {...props} showLoadingBar={isRequesting} />;
};
