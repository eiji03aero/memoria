import { ValidationResponseBody } from '@/domain/common/interfaces';

type Params = {
  responseBody: any;
};

export type ParsedValidationResponse = ReturnType<
  typeof useParseValidationResponse
>;

export const useParseValidationResponse = ({ responseBody }: Params) => {
  const isValidationResponse = typeof responseBody?.validation === 'object';
  let validationMessage: string | undefined;
  let key: string | undefined;
  let name: string | undefined;

  if (isValidationResponse) {
    const res: ValidationResponseBody = responseBody;
    key = res.validation.key;
    name = res.validation.name;
    switch (key) {
      case 'required':
        validationMessage = `${name} is required`;
        break;
      case 'invalid':
        validationMessage = `${name} is invalid`;
        break;
      case 'invalid-format':
        validationMessage = `${name} format is invalid`;
        break;
      case 'already-taken':
        validationMessage = `${name} is already taken`;
        break;
    }
  }

  return {
    isValidationResponse,
    validationMessage,
    key,
    name,
  };
};
