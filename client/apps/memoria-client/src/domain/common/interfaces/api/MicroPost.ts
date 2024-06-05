import { Medium } from '@/domain/common/interfaces/api/Medium';

export type MicroPost = {
  id: string;
  content: string;
  media: Medium[] | null;
};
