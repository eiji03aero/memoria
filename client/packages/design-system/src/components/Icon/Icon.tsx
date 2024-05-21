import { IconPropsWithoutChildren } from '@react-spectrum/icon';
import * as sstokens from '../../../styled-system/tokens';

import AddTo from '@spectrum-icons/workflow/AddTo';
import Article from '@spectrum-icons/workflow/Article';
import ImageAlbum from '@spectrum-icons/workflow/ImageAlbum';
import ImageCarousel from '@spectrum-icons/workflow/ImageCarousel';
import Settings from '@spectrum-icons/workflow/Settings';
import User from '@spectrum-icons/workflow/User';
import UserAdd from '@spectrum-icons/workflow/UserAdd';
import UserGroup from '@spectrum-icons/workflow/UserGroup';
import News from '@spectrum-icons/workflow/News';
import Add from '@spectrum-icons/workflow/Add';
import Back from '@spectrum-icons/workflow/Back';
import ChevronLeft from '@spectrum-icons/workflow/ChevronLeft';
import SelectAdd from '@spectrum-icons/workflow/SelectAdd';
import Cancel from '@spectrum-icons/workflow/Cancel';
import Delete from '@spectrum-icons/workflow/Delete';
import MoreVertical from '@spectrum-icons/workflow/MoreVertical';
import FolderAdd from '@spectrum-icons/workflow/FolderAdd';
import FolderRemove from '@spectrum-icons/workflow/FolderRemove';
import Remove from '@spectrum-icons/workflow/Remove';

// https://spectrum.adobe.com/page/icons/
const IconMap = {
  'add-to': AddTo,
  'image-album': ImageAlbum,
  'image-carousel': ImageCarousel,
  add: Add,
  article: Article,
  settings: Settings,
  user: User,
  'user-add': UserAdd,
  'user-group': UserGroup,
  news: News,
  back: Back,
  'chevron-left': ChevronLeft,
  'select-add': SelectAdd,
  cancel: Cancel,
  delete: Delete,
  'more-vertical': MoreVertical,
  'folder-add': FolderAdd,
  'folder-remove': FolderRemove,
  remove: Remove,
} as const;

export type IconName = keyof typeof IconMap;
export const IconNames = Object.keys(IconMap) as IconName[];

export type IconSize = IconPropsWithoutChildren['size'];

type Props = Omit<IconPropsWithoutChildren, 'color'> & {
  name: IconName;
  color: sstokens.ColorToken;
};

export const Icon = ({ name, color, ...rest }: Props) => {
  const Component = IconMap[name];
  const style = {
    color: sstokens.token.var(`colors.${color}`),
  };
  return <Component UNSAFE_style={style} {...rest} />;
};
