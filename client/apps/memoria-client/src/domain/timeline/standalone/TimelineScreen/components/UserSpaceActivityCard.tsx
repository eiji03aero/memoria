import * as React from 'react';
import { css } from '@/../styled-system/css';

import { useDateFormat } from '@/modules/hooks/useDateFormat';
import { UserSpaceActivity } from '@/domain/common/interfaces/api';

type Props = {
  userSpaceActivity: UserSpaceActivity;
};

export const UserSpaceActivityCard = ({ userSpaceActivity }: Props) => {
  const dfmt = useDateFormat();

  const { message } = React.useMemo(() => {
    if (userSpaceActivity.type === 'invite-user-joined') {
      return { message: `User ${userSpaceActivity.data.user_id} has joined!` };
    }
    if (userSpaceActivity.type === 'user-uploaded-media') {
      return { message: `User ${userSpaceActivity.data.user_id} uploaded media!` };
    }

    throw new Error(`UserSpaceActivityCard does not support type: ${userSpaceActivity}`);
  }, [userSpaceActivity]);

  return (
    <div className={Styles.card}>
      <div className={Styles.date}>{dfmt.pf.fullDateDOW(userSpaceActivity.created_at)}</div>
      <div className={Styles.message}>{message}</div>
    </div>
  );
};

const Styles = {
  card: css({
    p: '0.5rem',
    m: '0.5rem',
    bg: 'gray.100',
    borderRadius: 'md',
  }),
  date: css({
    fontSize: '0.75rem',
    color: 'gray.800',
  }),
  message: css({
    fontSize: 'md',
    color: 'black',
  }),
};
