import * as React from 'react';
import { useButton } from 'react-aria';

import * as styles from '../../styles';

import { Icon, IconName } from '../Icon';
import { Flex } from '../Flex';
import { View } from '../View';
import { CustomText } from '../CustomText';

type Props = {
  label: React.ReactNode;
  icon: IconName;
  active?: boolean;
  onPress?: () => void;
};

export const Item = ({ label, icon, active = false, onPress }: Props) => {
  const ref = React.useRef<HTMLButtonElement>(null);
  const { buttonProps } = useButton({ onPress }, ref);

  const contentColor = active ? 'yellow-500' : 'gray-600';

  return (
    <button {...buttonProps} className={styles.classnames(styles.reset.button)}>
      <View width="size-600" height="size-500" borderRadius="regular">
        <Flex
          width="100%"
          height="100%"
          direction="column"
          justifyContent="center"
          alignItems="center"
          gap="size-10"
        >
          <Icon size="S" color={contentColor} name={icon} />
          <CustomText size={50} color={contentColor}>
            {label}
          </CustomText>
        </Flex>
      </View>
    </button>
  );
};
