import { Thread } from '@/domain/common/interfaces/api/Thread';
import { User } from '@/domain/common/interfaces/api/User';

export type TimelinePost = {
  id: string;
  user: User;
  thread: Thread;
  created_at: string;
};
