import { ColorValue } from '@react-types/shared';

import * as styles from '../../styles';

const FontSizes = {
  50: { pc: 11, mobile: 13 },
  75: { pc: 12, mobile: 15 },
  100: { pc: 14, mobile: 17 },
  200: { pc: 16, mobile: 19 },
  300: { pc: 18, mobile: 22 },
  400: { pc: 20, mobile: 24 },
  500: { pc: 22, mobile: 27 },
  600: { pc: 25, mobile: 31 },
  700: { pc: 28, mobile: 34 },
  800: { pc: 32, mobile: 39 },
  900: { pc: 36, mobile: 44 },
  1000: { pc: 40, mobile: 49 },
  1100: { pc: 45, mobile: 55 },
  1200: { pc: 50, mobile: 62 },
  1300: { pc: 60, mobile: 70 },
} as const;

type FontSizeKey = keyof typeof FontSizes;

type Props = {
  as?: 'span' | 'p';
  size: FontSizeKey;
  color: ColorValue;
  className?: string;
  children: React.ReactNode;
};

export const CustomText = ({
  as: asProp = 'span',
  size,
  color,
  className,
  ...rest
}: Props) => {
  const {
    initialMatches: { mobile: isMobile },
  } = styles.useMediaQueries();

  const Component = asProp;
  const sizeConfig = FontSizes[size];
  const fontSize = isMobile ? sizeConfig.mobile : sizeConfig.pc;

  const style = {
    color: styles.CSSVar.color(color),
    fontSize,
  };
  return <Component className={className} style={style} {...rest} />;
};
