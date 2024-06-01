'use client';

import { useForm } from '@tanstack/react-form';
import { zodValidator } from '@tanstack/zod-form-adapter';
import { z } from 'zod';
import { Form, TextField, Button, TightLayoutCard } from '@repo/design-system';
import { css } from '@/../styled-system/css';

import { useInviteUser } from '@/domain/account/hooks/useInviteUser';
import { useParseValidationResponse } from '@/domain/common/hooks/useParseValidationResponse';
import { useValidationErrorProps } from '@/domain/common/hooks/useValidationErrorProps';

export default function InviteUser() {
  const { inviteUser, errorResponseBody } = useInviteUser();
  const parsedValidationResponse = useParseValidationResponse({
    responseBody: errorResponseBody,
  });
  const { buildErrorMessageProps } = useValidationErrorProps({
    parsedValidationResponse,
  });

  const form = useForm({
    defaultValues: {
      email: '',
    },
    onSubmit: async ({ value }) => {
      inviteUser(value);
    },
    validatorAdapter: zodValidator,
  });

  return (
    <TightLayoutCard.Background>
      <TightLayoutCard className={Styles.card}>
        <p className={Styles.instruction}>
          Let's invite user to this user space.
          <br />
        </p>
        <Form
          onSubmit={e => {
            e.preventDefault();
            form.handleSubmit();
          }}
        >
          <form.Field
            name="email"
            validators={{
              onChange: z.string().email(),
            }}
            children={field => (
              <TextField
                label="Email"
                name={field.name}
                value={field.state.value}
                {...buildErrorMessageProps({ field })}
                onChange={value => field.handleChange(value)}
              />
            )}
          />

          <div style={{ marginBottom: '2rem' }} />

          <form.Subscribe
            selector={state => [state.canSubmit, state.isSubmitting]}
            children={([canSubmit, isSubmitting]) => (
              <Button
                variant="primary"
                type="submit"
                isDisabled={!canSubmit || isSubmitting}
              >
                Submit
              </Button>
            )}
          />
        </Form>
      </TightLayoutCard>
    </TightLayoutCard.Background>
  );
}

const Styles = {
  card: css({
    padding: '0.5rem',
  }),
  instruction: css({
    fontSize: 'sm',
  }),
};