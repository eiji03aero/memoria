'use client';

import Image from 'next/image';
import { useSearchParams } from 'next/navigation';
import { useForm } from '@tanstack/react-form';
import { zodValidator } from '@tanstack/zod-form-adapter';
import { z } from 'zod';
import { Form, TextField, Button } from '@repo/design-system';
import { css } from '@/../styled-system/css';

import { useInviteUserConfirm } from '@/domain/account/hooks/useInviteUserConfirm';
import { useParseValidationResponse } from '@/domain/common/hooks/useParseValidationResponse';
import { useValidationErrorProps } from '@/domain/common/hooks/useValidationErrorProps';

export default function InviteUserConfirm() {
  const searchParams = useSearchParams();
  const invitationID = searchParams.get('id');

  const { inviteUserConfirm, errorResponseBody } = useInviteUserConfirm();
  const parsedValidationResponse = useParseValidationResponse({
    responseBody: errorResponseBody,
  });
  const { buildErrorMessageProps } = useValidationErrorProps({
    parsedValidationResponse,
  });

  const form = useForm({
    defaultValues: {
      name: '',
      password: '',
    },
    onSubmit: async ({ value }) => {
      if (!invitationID) {
        return;
      }

      inviteUserConfirm({
        name: value.name,
        password: value.password,
        invitationID,
      });
    },
    validatorAdapter: zodValidator,
  });

  return (
    <div className={Styles.container}>
      <Image
        src="/images/Logo-horizontal-black.png"
        width={300}
        height={72}
        alt="service logo"
        className={Styles.logo}
      />
      <div className={Styles.card}>
        <p className={Styles.instruction}>
          You are invited to memoria!
          <br />
          Please fill out the following form to join us :)
        </p>
        <Form
          onSubmit={e => {
            e.preventDefault();
            form.handleSubmit();
          }}
        >
          <form.Field
            name="name"
            validators={{
              onChange: z.string(),
            }}
            children={field => (
              <TextField
                label="Name"
                name={field.name}
                value={field.state.value}
                {...buildErrorMessageProps({ field })}
                onChange={value => field.handleChange(value)}
              />
            )}
          />
          <form.Field
            name="password"
            validators={{
              onChange: z.string(),
            }}
            children={field => (
              <TextField
                label="Password"
                type="password"
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
      </div>
    </div>
  );
}

const Styles = {
  container: css({
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    paddingTop: '4rem',
  }),
  logo: css({
    marginBottom: '2rem',
  }),
  instruction: css({
    marginBottom: '0.5rem',
  }),
  card: css({
    background: 'gray.200',
    borderRadius: 'lg',
    maxWidth: '320px',
    padding: '1rem',
  }),
};
