import * as React from 'react';
import { FieldApi } from '@tanstack/react-form';

import { ParsedValidationResponse } from '@/domain/common/hooks/useParseValidationResponse';

type Params = {
  parsedValidationResponse: ParsedValidationResponse;
};

export const useValidationErrorProps = ({
  parsedValidationResponse,
}: Params) => {
  const getErrorMessage = React.useCallback(
    (params: { field: FieldApi<unknown, any>; name?: string }) => {
      const name = params.name ?? params.field.name;
      if (name === parsedValidationResponse.name) {
        return parsedValidationResponse.validationMessage;
      }

      const { errors } = params.field.state.meta;
      if (errors.length > 0) {
        return errors.join(',');
      }

      return undefined;
    },
    [parsedValidationResponse],
  );

  const buildErrorMessageProps = React.useCallback(
    // Ah let us fuck with FieldApi type params
    (params: { field: FieldApi<any, any, any, any, any>; name?: string }) => {
      const { isDirty } = params.field.state.meta;

      const errorMessage = getErrorMessage(params);

      let validationState: 'valid' | 'invalid' | undefined;
      if (isDirty) {
        validationState = errorMessage
          ? ('invalid' as const)
          : ('valid' as const);
      }

      return {
        validationState,
        errorMessage,
      };
    },
    [getErrorMessage],
  );

  return {
    buildErrorMessageProps,
  };
};
