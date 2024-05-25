import * as React from 'react';
import { css } from '@/../styled-system/css';

import { useMediumGetPage } from '@/domain/media/hooks/useMediumGetPage';
import { useMedia } from '@/domain/media/hooks/useMedia';
import { MediaSlider } from '@/domain/media/standalone/MediaGalleryScreen/components/MediaSlider';
import { ScreenTools } from '@/domain/media/standalone/MediaGalleryScreen/components/ScreenTools';

type Props = {
  albumID?: string;
  initialMediumID: string;
  onBack: () => void;
};

export const MediaGalleryScreen = ({ albumID, initialMediumID, onBack }: Props) => {
  const [showTools, setShowTools] = React.useState(false);

  const { pagination, isFetched } = useMediumGetPage({ albumID, mediumID: initialMediumID });
  const { media, isFetching, fetchNextPage, fetchPreviousPage, hasNextPage, hasPreviousPage } =
    useMedia({
      albumID,
      enabled: isFetched,
      initialPage: pagination?.current_page,
    });
  const [cursor, setCursor] = React.useState(initialMediumID);
  const cursorIdx = media?.findIndex(m => m.id === cursor);

  const onNext = () => {
    if (cursorIdx === undefined || isFetching) {
      return;
    }

    if (hasNextPage) {
      fetchNextPage();
      return;
    }
  };

  const onPrev = () => {
    if (cursorIdx === undefined || isFetching) {
      return;
    }

    if (hasPreviousPage) {
      fetchNextPage();
      return;
    }
  };

  const slides = React.useMemo(() => {
    if (!media) {
      return [];
    }

    return media.map(m => ({
      type: m.type,
      src: m.original_url,
    }));
  }, [media]);

  const handleClickScreen = () => {
    setShowTools(prev => !prev);
  };

  return (
    <div className={Styles.screen} onClick={handleClickScreen}>
      {typeof cursorIdx !== 'undefined' && (
        <MediaSlider
          slides={slides}
          initialSlide={cursorIdx}
          onEdge={dir => {
            if (dir === 'right') {
              onPrev();
            }
            if (dir === 'left') {
              onNext();
            }
            console.log('gonna dir', dir);
          }}
        />
      )}

      <ScreenTools show={showTools} onBack={onBack} />
    </div>
  );
};

const Styles = {
  screen: css({
    height: '100%',
    bg: 'black',
  }),
};
