import { TimelinePost } from '@/domain/common/interfaces/api/TimelinePost';
import { UserSpaceActivity } from '@/domain/common/interfaces/api/UserSpaceActivity';

type Base = {
  id: string;
};

type TimelineUnit_USA = Base & {
  type: 'user-space-activity';
  data: UserSpaceActivity;
};

type TimelineUnit_TP = Base & {
  type: 'timeline-post';
  data: TimelinePost;
};

export type TimelineUnit = TimelineUnit_USA | TimelineUnit_TP;
