import {
  TextField as BaseTextField,
  SpectrumTextFieldProps,
} from '@adobe/react-spectrum';

import * as styles from '../../styles';
import { css } from '../../../styled-system/css';

type Props = Omit<SpectrumTextFieldProps, 'UNSAFE_className'> & {
  className?: string;
};

export const TextField = ({ className, ...rest }: Props) => {
  const classNames = styles.classnames(Styles.field, className);
  return <BaseTextField UNSAFE_className={classNames} {...rest} />;
};

const Styles = {
  field: css({
    '& svg': {
      boxSizing: 'content-box',
    },
  }),
};
