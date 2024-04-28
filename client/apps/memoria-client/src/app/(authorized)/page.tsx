'use client';
import { useAppData } from '@/domain/common/hooks/useAppData';

export default function Top() {
  const { state } = useAppData();
  console.log(state.user.id);

  return (
    <div>
      <h1>Top page will come here!</h1>

      <p>user: {JSON.stringify(state.user)}</p>
      <p>user space: {JSON.stringify(state.userSpace)}</p>
    </div>
  );
}
