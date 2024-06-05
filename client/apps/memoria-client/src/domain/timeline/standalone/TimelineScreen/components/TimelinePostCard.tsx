import { TightLayoutCard, Icon } from '@repo/design-system';
import { css } from '@/../styled-system/css';

import { useDateFormat } from '@/modules/hooks/useDateFormat';
import { TimelinePost } from '@/domain/common/interfaces/api/TimelinePost';

type Props = {
  timelinePost: TimelinePost;
};

export const TimelinePostCard = ({ timelinePost }: Props) => {
  const dfmt = useDateFormat();
  return (
    <TightLayoutCard className={Styles.card}>
      <div className={Styles.header.box}>
        <div className={Styles.header.avatar}>
          <Icon name="user" color="gray.400" size="S" />
        </div>
        <div className={Styles.header.main}>
          <div className={Styles.header.name}>{timelinePost.user.name}</div>
          <div className={Styles.header.date}>{dfmt.pf.fullDateDOW(timelinePost.created_at)}</div>
        </div>
      </div>
      <div className={Styles.message}>{timelinePost.thread.micro_posts[0]?.content}</div>
      <div className={Styles.mediaArea}>
        {timelinePost.thread.micro_posts[0]?.media?.slice(0, 3).map(m => (
          <div className={Styles.media}>
            <img src={m.tn_240_url} />
          </div>
        ))}
      </div>
    </TightLayoutCard>
  );
};

const Styles = {
  card: css({
    padding: '0.5rem',
  }),
  header: {
    box: css({
      display: 'flex',
    }),
    avatar: css({
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'center',
      w: '2rem',
      h: '2rem',
      borderRadius: '50%',
      bg: 'gray.100',
      mr: '1rem',
    }),
    main: css({
      flex: 1,
    }),
    name: css({
      fontSize: '1rem',
      color: 'gray.800',
      lineHeight: 1,
      mb: '0.25rem',
    }),
    date: css({
      fontSize: '0.75rem',
      color: 'gray.800',
      lineHeight: 1,
      mb: '0.5rem',
    }),
  },
  message: css({
    fontSize: 'md',
    color: 'black',
    mb: '0.5rem',
  }),
  mediaArea: css({
    display: 'flex',
    gap: '0.5rem',
  }),
  media: css({
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    w: 100,
    h: 100,
    '& > img': {
      width: '100%',
      maxHeight: '100%',
    },
  }),
};
