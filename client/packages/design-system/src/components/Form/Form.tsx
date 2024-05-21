import { Form as BaseForm, SpectrumFormProps } from '@adobe/react-spectrum';
import { css } from '../../../styled-system/css';

import * as styles from '../../styles';

export const Form = ({ UNSAFE_className, ...rest }: SpectrumFormProps) => {
  return (
    <BaseForm
      UNSAFE_className={styles.classnames(
        css({ '& > *': { mt: '0 !important' } }),
        UNSAFE_className,
      )}
      {...rest}
    />
  );
};
