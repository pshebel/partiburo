import { useAppForm } from '../../hooks/form.tsx'
import { loginFormOptions } from './guest-form-options.tsx'
import {  useMutation } from '@tanstack/react-query';
import { useNavigate } from 'react-router-dom'
import { getGuest } from '../../hooks/identity'

export const Guest = () => {
  const navigate = useNavigate()

  const form = useAppForm({
    ...loginFormOptions,
    onSubmit: async ({ formApi, value }) => {
      await saveUserMutation.mutateAsync(value)

      // Reset the form to start-over with a clean state
      formApi.reset()
    },
  })

  const saveUserMutation = useMutation({
    mutationFn: async (req: { status: string, phone: string }) => {
        const guest_id = getGuest()
        const body = {
          id: guest_id,
          status: req.status,
          phone: req.phone,
        }
        const response = await fetch(`${import.meta.env.VITE_API_URL}/guest`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                },
                body: JSON.stringify(body),
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
      <h1>Update RSVP or Phone Number</h1>
      <form.AppField
        name="phone"
        children={(field) => <field.TextField label="Phone" />}
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
