export type ValidationKey = 'required' | 'invalid-format' | 'already-taken';

export type ValidationResponseBody = {
  validation: {
    key: ValidationKey;
    name: string;
  };
};
