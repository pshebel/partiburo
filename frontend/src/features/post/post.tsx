import { useAppForm } from '../../hooks/form.tsx'
import { loginFormOptions } from './login-form-options.tsx'
import {  useMutation } from '@tanstack/react-query';

import { createGuest } from '../../hooks/identity';
import { Response } from '../../interfaces/response.js'


interface LoginProps {
    onLoginSuccess: (id: string) => void;
}

export const CreatePost = ({ onLoginSuccess }: LoginProps) => {
  const form = useAppForm({
    ...loginFormOptions,
    onSubmit: async ({ formApi, value }) => {
      await saveUserMutation.mutateAsync(value)

      // Reset the form to start-over with a clean state
      formApi.reset()
    },
  })

  const saveUserMutation = useMutation({
    mutationFn: async (req: { name: string; status: string }) => {
    //   const response = await fetch(`${process.env.EXPO_PUBLIC_API_URL}`, {
        const response = await fetch('http://localhost:4000/guest', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                },
                body: JSON.stringify(req),
        });
        return response.json() as Promise<Response>;
    },
    onSuccess: (data: Response) => {
        if (data.Code === 200) {
            createGuest(data.Message)
            onLoginSuccess(data.Message);
        } else {
            window.confirm(data.Message)
        }
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
      <h1>Identify Yourself</h1>
      <form.AppField
        name="name"
        children={(field) => <field.TextField label="Name" />}
      />
      <form.Field
        name="status" // This is the field name in defaultValues
        children={(field) => (
          <div>
            <label>Attending:</label>
            <div>
              <input
                type="radio"
                id="GOING"
                name={field.name} // Use the field name for grouping
                value="GOING"
                checked={field.state.value === 'GOING'} // Bind checked state
                onChange={(e) => field.handleChange(e.target.value)} // Update form state on change
              />
              <label htmlFor="GOING">Going</label>
            </div>

            <div>
              <input
                type="radio"
                id="MAYBE"
                name={field.name}
                value="MAYBE"
                checked={field.state.value === 'MAYBE'}
                onChange={(e) => field.handleChange(e.target.value)}
              />
              <label htmlFor="MAYBE">Maybe</label>
            </div>
            <div>
              <input
                type="radio"
                id="NOT_GOING"
                name={field.name}
                value="NOT_GOING"
                checked={field.state.value === 'NOT_GOING'}
                onChange={(e) => field.handleChange(e.target.value)}
              />
              <label htmlFor="NOT_GOING">Not Going</label>
            </div>

            {/* Display potential errors */}
            {field.state.meta.errors ? (
                <em>{field.state.meta.errors.join(', ')}</em>
            ) : null}
          </div>
        )}
      />

      <form.AppForm>
        <form.SubscribeButton label="Submit" />
      </form.AppForm>
    </form>
  )
}
