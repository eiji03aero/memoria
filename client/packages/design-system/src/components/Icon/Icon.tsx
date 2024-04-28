import { IconPropsWithoutChildren } from '@react-spectrum/icon';
import { ColorValue } from '@react-types/shared';

import ImageAlbum from '@spectrum-icons/workflow/ImageAlbum';
import AddTo from '@spectrum-icons/workflow/AddTo';
import Add from '@spectrum-icons/workflow/Add';
import Settings from '@spectrum-icons/workflow/Settings';

import * as styles from '../../styles';

// https://spectrum.adobe.com/page/icons/
const IconMap = {
  'image-album': ImageAlbum,
  'add-to': AddTo,
  add: Add,
  settings: Settings,
} as const;

export type IconName = keyof typeof IconMap;
export const IconNames = Object.keys(IconMap) as IconName[];

type Props = Omit<IconPropsWithoutChildren, 'color'> & {
  name: IconName;
  color: ColorValue;
};

export const Icon = ({ name, color, ...rest }: Props) => {
  const Component = IconMap[name];
  const style = {
    color: styles.CSSVar.color(color),
  };
  return <Component UNSAFE_style={style} {...rest} />;
};
