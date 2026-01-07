import { useAppForm } from '../../hooks/form.tsx'
import { postFormOptions } from './post-form-options.tsx'
import {  useMutation } from '@tanstack/react-query';
import { getGuest } from '../../hooks/identity'
import { Response } from '../../interfaces/response.js'
import { useNavigate } from 'react-router-dom'

export const Post = () => {
  const navigate = useNavigate()
  const guest_id = getGuest()
  const form = useAppForm({
    ...postFormOptions,
    onSubmit: async ({ formApi, value }) => {
      await saveUserMutation.mutateAsync(value)

      // Reset the form to start-over with a clean state
      formApi.reset()
    },
  })

  const saveUserMutation = useMutation({
    mutationFn: async (req: { body: string;}) => {
        const r = {
          id: guest_id,
          body: req.body
        }
        const response = await fetch(`${import.meta.env.VITE_API_URL}/post`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                },
                body: JSON.stringify(r),
        });
        return response.json() as Promise<Response>;
    },
    onSuccess: () => {
      navigate('/')
    },
    onError: (err: any) => {
        window.confirm(err)
    },
  })

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault()
        form.handleSubmit()
      }}
    >
      <h1>Create Post</h1>
      <form.AppField
        name="body"
        children={(field) => <field.TextArea label="Body" />}
      />

      <form.AppForm>
        <form.SubscribeButton label="Submit" />
      </form.AppForm>
    </form>
  )
}
