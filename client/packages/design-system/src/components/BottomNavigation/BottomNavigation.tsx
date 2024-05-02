import { Item } from './Item';
import { Flex } from '../Flex';
import { View } from '../View';

export const BottomNavigation = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  return (
    <View
      elementType="nav"
      paddingY="size-100"
      borderTopWidth="thin"
      borderTopColor="dark"
      backgroundColor="gray-50"
    >
      <Flex justifyContent="space-around">{children}</Flex>
    </View>
  );
};

BottomNavigation.Item = Item;
