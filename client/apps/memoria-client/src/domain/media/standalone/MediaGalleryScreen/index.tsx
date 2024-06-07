import * as React from 'react';
import type { Slider as SliderT } from '@repo/design-system';
import { css } from '@/../styled-system/css';

import { useMediumGetPage } from '@/domain/media/hooks/useMediumGetPage';
import { useMedia } from '@/domain/media/hooks/useMedia';
import { MediaSlider } from '@/domain/media/standalone/MediaGalleryScreen/components/MediaSlider';
import { ScreenTools } from '@/domain/media/standalone/MediaGalleryScreen/components/ScreenTools';
import { usePreviousValue } from '@/modules/hooks/usePreviousValue';

type Props = {
  albumID?: string;
  initialMediumID: string;
  onBack: () => void;
};

export const MediaGalleryScreen = ({ albumID, initialMediumID, onBack }: Props) => {
  const sliderRef = React.useRef(null) as React.MutableRefObject<SliderT | null>;
  const [showTools, setShowTools] = React.useState(true);

  const { pagination, isFetched } = useMediumGetPage({ albumID, mediumID: initialMediumID });
  const { media, isFetching, fetchNextPage, fetchPreviousPage, hasNextPage, hasPreviousPage } =
    useMedia({
      albumID,
      enabled: isFetched,
      initialPosition: pagination?.current_page,
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
      fetchPreviousPage();
      return;
    }
  };

  const slides = React.useMemo(() => {
    if (!media) {
      return [];
    }

    return media.map(m => ({
      type: m.type,
      id: m.id,
      src: m.original_url,
    }));
  }, [media]);

  const handleClickScreen = () => {
    setShowTools(prev => !prev);
  };

  const { previousValue: prevci } = usePreviousValue(cursorIdx);
  React.useEffect(() => {
    if (cursorIdx === undefined || prevci === undefined || cursorIdx === prevci) {
      return;
    }

    if (cursorIdx > prevci) {
      window.setTimeout(() => {
        sliderRef.current?.slickGoTo(cursorIdx - 1, false);
      }, 100);
    }
  }, [cursorIdx, prevci]);

  return (
    <div className={Styles.screen} onClick={handleClickScreen}>
      {typeof cursorIdx !== 'undefined' && (
        <MediaSlider
          sliderRef={sliderRef}
          slides={slides}
          initialSlide={cursorIdx}
          showTools={showTools}
          afterChange={idx => {
            const s = slides[idx];
            if (s) {
              setCursor(s.id);
            }
          }}
          onEdge={dir => {
            if (dir === 'right') {
              onPrev();
            }
            if (dir === 'left') {
              onNext();
            }
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
