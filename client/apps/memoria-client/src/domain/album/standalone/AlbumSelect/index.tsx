import * as React from 'react';
import { Select } from '@repo/design-system';

import { useAlbums } from '@/domain/album/hooks/useAlbums';

type Props = {
  onChange: (values: readonly { label: string; value: string }[]) => void;
};

export const AlbumSelect = ({ onChange }: Props) => {
  const { albums, isFetching } = useAlbums();
  const options = React.useMemo(() => {
    if (!albums) {
      return [];
    }

    return albums.map(album => ({
      label: album.name,
      value: album.id,
    }));
  }, [albums]);

  return (
    <Select<{ label: string; value: string }, true>
      isMulti={true}
      isLoading={isFetching}
      onChange={values => onChange(values)}
      options={options}
      closeMenuOnSelect={false}
      blurInputOnSelect={false}
    />
  );
};
