'use client';

import { useTranslation } from 'react-i18next';
import Image from 'next/image';
import { useForm } from '@tanstack/react-form';
import { zodValidator } from '@tanstack/zod-form-adapter';
import { z } from 'zod';
import { Form, TextField, Button } from '@repo/design-system';
import { css } from '@/../styled-system/css';

import { useSignup } from '@/domain/account/hooks/useSignup';
import { useParseValidationResponse } from '@/domain/common/hooks/useParseValidationResponse';
import { useValidationErrorProps } from '@/domain/common/hooks/useValidationErrorProps';

export default function Signup() {
  const { t } = useTranslation();
  const { signup, errorResponseBody } = useSignup();
  const parsedValidationResponse = useParseValidationResponse({
    responseBody: errorResponseBody,
  });
  const { buildErrorMessageProps } = useValidationErrorProps({
    parsedValidationResponse,
  });

  const form = useForm({
    defaultValues: {
      name: '',
      email: '',
      userSpaceName: '',
      password: '',
    },
    onSubmit: async ({ value }) => {
      signup(value);
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
          {t('p.signup.headings')}
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
            name="name"
            validators={{
              onChange: z.string(),
            }}
            children={field => (
              <TextField
                label={t('w.name')}
                name={field.name}
                value={field.state.value}
                {...buildErrorMessageProps({ field })}
                onChange={value => field.handleChange(value)}
              />
            )}
          />
          <form.Field
            name="userSpaceName"
            validators={{
              onChange: z.string(),
            }}
            children={field => (
              <TextField
                label={t('w.user-space-name')}
                name={field.name}
                value={field.state.value}
                {...buildErrorMessageProps({ field, name: 'user_space_name' })}
                onChange={value => field.handleChange(value)}
              />
            )}
          />
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
              <div
                className={css({
                  pt: '0.5rem',
                })}
              >
                <Button variant="primary" type="submit" isDisabled={!canSubmit || isSubmitting}>
                  {t('w.submit')}
                </Button>
              </div>
            )}
          />
        </Form>
      </div>
    </>
  );
}
