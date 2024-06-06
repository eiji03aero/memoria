'use client';

import { useTranslation } from 'react-i18next';
import Image from 'next/image';
import { useForm } from '@tanstack/react-form';
import { zodValidator } from '@tanstack/zod-form-adapter';
import { z } from 'zod';
import { TextField, Button, Form } from '@repo/design-system';
import { css } from '@/../styled-system/css';

import { useLogin } from '@/domain/account/hooks/useLogin';
import { useParseValidationResponse } from '@/domain/common/hooks/useParseValidationResponse';
import { useValidationErrorProps } from '@/domain/common/hooks/useValidationErrorProps';

export default function Login() {
  const { t } = useTranslation();
  const { login, errorResponseBody } = useLogin();
  const parsedValidationResponse = useParseValidationResponse({
    responseBody: errorResponseBody,
  });
  const { buildErrorMessageProps } = useValidationErrorProps({
    parsedValidationResponse,
  });

  const form = useForm({
    defaultValues: {
      email: '',
      password: '',
    },
    onSubmit: async ({ value }) => {
      login(value);
    },
    validatorAdapter: zodValidator,
  });

  return (
    <>
      <div
        className={css({
          display: 'flex',
          alignContent: 'cetner',
          justifyContent: 'center',
          width: '100%',
          marginY: '3rem',
        })}
      >
        <Image src="/images/Logo-horizontal-black.png" width={300} height={72} alt="service logo" />
      </div>

      <div
        className={css({
          display: 'flex',
          justifyContent: 'center',
          marginBottom: '1rem',
        })}
      >
        <p
          className={css({
            textAlign: 'center',
            fontSize: '1rem',
            color: 'black',
            whiteSpace: 'pre-wrap',
          })}
        >
          {t('p.login.headings')}
        </p>
      </div>

      <div
        className={css({
          marginX: 'auto',
          maxWidth: '300px',
          backgroundColor: 'gray.200',
          borderRadius: 'lg',
          padding: '1rem',
        })}
      >
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
                label={t('w.email')}
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
                label={t('w.password')}
                type="password"
                name={field.name}
                value={field.state.value}
                {...buildErrorMessageProps({ field })}
                onChange={value => field.handleChange(value)}
              />
            )}
          />

          <form.Subscribe
            selector={state => [state.canSubmit, state.isSubmitting]}
            children={([canSubmit, isSubmitting]) => (
              <Button variant="primary" type="submit" isDisabled={!canSubmit || isSubmitting}>
                {t('w.submit')}
              </Button>
            )}
          />
        </Form>
      </div>
    </>
  );
}
