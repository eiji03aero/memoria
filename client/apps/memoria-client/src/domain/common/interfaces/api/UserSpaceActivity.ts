type Base = {
  id: string;
  created_at: string;
};

export type UserSpaceActivity_InviteUserJoined = Base & {
  type: 'invite-user-joined';
  data: {
    user_id: string;
  };
};

export type UserSpaceActivity_UserUploadedMedia = Base & {
  type: 'user-uploaded-media';
  data: {
    user_id: string;
    medium_ids: string[];
  };
};

export type UserSpaceActivity =
  | UserSpaceActivity_InviteUserJoined
  | UserSpaceActivity_UserUploadedMedia;
