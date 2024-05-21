import * as React from 'react';

import { MediaGrid } from '@repo/design-system';
import { Medium } from '@/domain/common/interfaces/api';

type Props = {
  medium: Medium;
  selected?: boolean;
  onPress?: () => void;
};

export const MediaGridTile = ({ medium, selected, onPress }: Props) => {
  const [srcIdx, setSrcIdx] = React.useState(0);

  const src = React.useMemo(() => {
    switch (srcIdx) {
      case 0:
        return medium.tn_240_url;
      case 1:
        return medium.original_url;
      default:
        return undefined;
    }
  }, [medium, srcIdx]);

  const handleError = React.useCallback(() => {
    setSrcIdx(prev => prev + 1);
  }, []);

  return (
    <MediaGrid.Tile src={src} selected={selected} onPress={onPress} onImgError={handleError} />
  );
};
