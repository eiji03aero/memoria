import { useTranslation } from 'react-i18next';

import { ValidationResponseBody } from '@/domain/common/interfaces';

type Params = {
  responseBody: any;
};

export type ParsedValidationResponse = ReturnType<typeof useParseValidationResponse>;

export const useParseValidationResponse = ({ responseBody }: Params) => {
  const { t } = useTranslation();
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
        validationMessage = t('v.data-is-required', { data: name });
        break;
      case 'invalid':
        validationMessage = t('v.data-is-invalid', { data: name });
        break;
      case 'invalid-format':
        validationMessage = t('v.data-format-is-invalid', { data: name });
        break;
      case 'already-taken':
        validationMessage = t('v.data-is-already-taken', { data: name });
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
