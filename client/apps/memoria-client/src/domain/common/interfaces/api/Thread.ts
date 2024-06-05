import { MicroPost } from '@/domain/common/interfaces/api/MicroPost';

export type Thread = {
  id: string;
  micro_posts: MicroPost[];
};
