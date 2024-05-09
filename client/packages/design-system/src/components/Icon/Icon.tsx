import { IconPropsWithoutChildren } from '@react-spectrum/icon';
import * as sstokens from '../../../styled-system/tokens';

import Add from '@spectrum-icons/workflow/Add';
import AddTo from '@spectrum-icons/workflow/AddTo';
import Article from '@spectrum-icons/workflow/Article';
import ImageAlbum from '@spectrum-icons/workflow/ImageAlbum';
import ImageCarousel from '@spectrum-icons/workflow/ImageCarousel';
import Settings from '@spectrum-icons/workflow/Settings';
import User from '@spectrum-icons/workflow/User';
import UserAdd from '@spectrum-icons/workflow/UserAdd';
import UserGroup from '@spectrum-icons/workflow/UserGroup';

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
} as const;

export type IconName = keyof typeof IconMap;
export const IconNames = Object.keys(IconMap) as IconName[];

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
