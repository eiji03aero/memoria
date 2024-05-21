'use client';

import { useForm } from '@tanstack/react-form';
import { zodValidator } from '@tanstack/zod-form-adapter';
import { z } from 'zod';
import { Form, Button, BottomDrawer, TextField } from '@repo/design-system';
import { css } from '@/../styled-system/css';

import { useCreateAlbum } from '@/domain/album/hooks/useCreateAlbum';
import { useParseValidationResponse } from '@/domain/common/hooks/useParseValidationResponse';
import { useValidationErrorProps } from '@/domain/common/hooks/useValidationErrorProps';

type Props = {
  show: boolean;
  onClose: () => void;
};

export const CreateAlbumDrawer = ({ show, onClose }: Props) => {
  const { createAlbum, errorResponseBody } = useCreateAlbum();
  const parsedValidationResponse = useParseValidationResponse({
    responseBody: errorResponseBody,
  });
  const { buildErrorMessageProps } = useValidationErrorProps({
    parsedValidationResponse,
  });

  const form = useForm({
    defaultValues: {
      name: '',
    },
    onSubmit: async ({ value }) => {
      createAlbum(value, {
        onSuccess: () => {
          onClose();
        },
      });
    },
    validatorAdapter: zodValidator,
  });

  return (
    <BottomDrawer show={show} onClose={onClose}>
      <Form
        onSubmit={e => {
          e.preventDefault();
          form.handleSubmit();
        }}
      >
        <BottomDrawer.Content>
          <BottomDrawer.Header onClose={onClose}>
            Create an album
          </BottomDrawer.Header>

          <BottomDrawer.Body className={css({ mb: '1rem' })}>
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
          </BottomDrawer.Body>

          <BottomDrawer.Footer>
            <Button variant="primary" type="submit">
              Save
            </Button>
          </BottomDrawer.Footer>
        </BottomDrawer.Content>
      </Form>
    </BottomDrawer>
  );
};
