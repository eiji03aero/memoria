import * as React from 'react';
import Image from 'next/image';
import { useTranslation } from 'react-i18next';
import { TightLayoutCard, Button, IconButton, LoadingBlock } from '@repo/design-system';
import { css } from '@/../styled-system/css';

import { Link } from '@/modules/components';

import { TopicHeader } from '@/domain/common/standalone/TopicHeader';
import { useTimelineUnits } from '@/domain/timeline/hooks/useTimelineUnits';
import { UserSpaceActivityCard } from '@/domain/timeline/standalone/TimelineScreen/components/UserSpaceActivityCard';
import { TimelinePostCard } from '@/domain/timeline/standalone/TimelineScreen/components/TimelinePostCard';
import { TimelinePostDrawer } from '@/domain/timeline/standalone/TimelinePostDrawer';

export const TimelineScreen = () => {
  const { t } = useTranslation();
  const [showCreatePost, setShowCreatePost] = React.useState(false);
  const { units, fetchNextPage, isFetching } = useTimelineUnits();

  return (
    <>
      <TopicHeader
        stickyTop
        label={
          <Link href="/timeline">
            <Image
              src="/images/Logo-horizontal-black.png"
              width={150}
              height={36}
              alt="service logo"
            />
          </Link>
        }
        trailing={
          <>
            <IconButton variant="elevated" iconName="add" onPress={() => setShowCreatePost(true)} />
          </>
        }
      />

      <TightLayoutCard.Background>
        {units?.map(unit => {
          if (unit.type === 'user-space-activity') {
            return <UserSpaceActivityCard key={unit.id} userSpaceActivity={unit.data} />;
          }
          if (unit.type === 'timeline-post') {
            return <TimelinePostCard key={unit.id} timelinePost={unit.data} />;
          }
        })}

        {isFetching && <LoadingBlock />}
        {!isFetching && (
          <div className={Styles.loadMoreBox}>
            <Button variant="primary" onPress={() => fetchNextPage()} width="100%">
              {t('w.load-more')}
            </Button>
          </div>
        )}
      </TightLayoutCard.Background>

      <TimelinePostDrawer show={showCreatePost} onClose={() => setShowCreatePost(false)} />
    </>
  );
};

const Styles = {
  loadMoreBox: css({
    display: 'flex',
    m: '0.5rem',
  }),
};
