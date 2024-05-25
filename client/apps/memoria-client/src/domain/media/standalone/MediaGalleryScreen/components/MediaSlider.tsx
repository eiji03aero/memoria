import * as React from 'react';
import { Slider } from '@repo/design-system';
import { css } from '@/../styled-system/css';

import { Slide } from '@/domain/media/standalone/MediaGalleryScreen/interfaces/Slide';
import { ImagePanel } from '@/domain/media/standalone/MediaGalleryScreen/components/ImagePanel';
import { VideoPanel } from '@/domain/media/standalone/MediaGalleryScreen/components/VideoPanel';

type Props = {
  slides: Slide[];
  initialSlide: number;
  onEdge: React.ComponentProps<typeof Slider>['onEdge'];
};

export const MediaSlider = ({ slides, initialSlide, onEdge }: Props) => {
  return (
    <Slider
      className={Styles.slider}
      lazyLoad="ondemand"
      slidesToShow={1}
      initialSlide={initialSlide}
      infinite={false}
      onEdge={onEdge}
    >
      {slides.map(slide => {
        if (slide.type === 'image') {
          return <ImagePanel key={slide.src} src={slide.src} />;
        }
        if (slide.type === 'video') {
          return <VideoPanel key={slide.src} src={slide.src} />;
        }

        throw new Error(`unknown slide type: ${slide.type}`);
      })}
    </Slider>
  );
};

const Styles = {
  slider: css({
    '& .slick-prev': {
      left: '1.5rem',
    },
    '& .slick-next': {
      right: '1.5rem',
    },
    '& .slick-next, .slick-prev': {
      zIndex: 1,
    },
  }),
};
