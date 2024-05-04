'use client';

import Image from 'next/image';
import { useForm } from '@tanstack/react-form';
import { zodValidator } from '@tanstack/zod-form-adapter';
import { z } from 'zod';
import {
  Flex,
  Form,
  TextField,
  View,
  CustomText,
  Button,
} from '@repo/design-system';

import { useLogin } from '@/domain/account/hooks/useLogin';
import { useParseValidationResponse } from '@/domain/common/hooks/useParseValidationResponse';
import { useValidationErrorProps } from '@/domain/common/hooks/useValidationErrorProps';

export default function Login() {
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
      <Flex
        width="100%"
        height="size-1000"
        marginTop="size-1000"
        marginBottom="size-500"
        alignItems="center"
        justifyContent="center"
      >
        <Image
          src="/images/Logo-horizontal-black.png"
          width={300}
          height={72}
          alt="service logo"
        />
      </Flex>
      <View marginBottom="size-500">
        <Flex justifyContent="center">
          <CustomText className="text-center" size={200} color="gray-900">
            Let us begin a long-lasting relationship
            <br />
            with your memories
          </CustomText>
        </Flex>
      </View>
      <View
        marginX="auto"
        maxWidth="size-3400"
        backgroundColor="gray-200"
        borderRadius="large"
        padding="size-200"
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
                label="Email"
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
                label="password"
                type="password"
                name={field.name}
                value={field.state.value}
                {...buildErrorMessageProps({ field })}
                onChange={value => field.handleChange(value)}
              />
            )}
          />

          <View marginBottom="size-200" />

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
      </View>
    </>
  );
}
