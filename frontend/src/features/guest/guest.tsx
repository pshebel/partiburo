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
    mutationFn: async (req: { status: string, email: string }) => {
        const guest_id = getGuest()
        const body = {
          id: guest_id,
          status: req.status,
          email: req.email,
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
    <div className="min-h-screen bg-gray-50 py-12 px-4 flex items-center justify-center">
      <form
        onSubmit={(e) => {
          e.preventDefault()
          form.handleSubmit()
        }}
        className="max-w-md w-full bg-white p-8 rounded-2xl shadow-sm border border-gray-100 space-y-8"
      >
        {/* Header section */}
        <div className="border-b pb-4">
          <h1 className="text-xs font-bold uppercase tracking-widest text-blue-600 mb-1">Preferences</h1>
          <h2 className="text-2xl font-extrabold text-gray-900 leading-tight">Update RSVP or Email</h2>
        </div>

        <div className="space-y-6">
          {/* Email Field - Uses your updated TextField internally */}
          <form.AppField
            name="email"
            children={(field) => <field.TextField label="Email (optional)" placeholder="Enter an email if you wish to receive alerts" />}
          />

          {/* Status Radio Group */}
          <form.Field
            name="status"
            children={(field) => (
              <div className="space-y-3">
                <label className="text-sm font-semibold text-gray-700">Will you be attending?</label>
                <div className="grid grid-cols-1 gap-2">
                  {[
                    { id: 'GOING', label: 'Going' },
                    { id: 'MAYBE', label: 'Maybe' },
                    { id: 'NOT_GOING', label: 'Not Going' }
                  ].map((option) => (
                    <label 
                      key={option.id}
                      className={`
                        flex items-center px-4 py-3 rounded-xl border cursor-pointer transition-all
                        ${field.state.value === option.id 
                          ? 'bg-blue-50 border-blue-600 ring-1 ring-blue-600' 
                          : 'bg-white border-gray-200 hover:bg-gray-50'
                        }
                      `}
                    >
                      <input
                        type="radio"
                        className="w-4 h-4 text-blue-600 border-gray-300 focus:ring-blue-500"
                        name={field.name}
                        value={option.id}
                        checked={field.state.value === option.id}
                        onChange={(e) => field.handleChange(e.target.value)}
                      />
                      <span className={`ml-3 text-sm font-bold ${field.state.value === option.id ? 'text-blue-700' : 'text-gray-600'}`}>
                        {option.label}
                      </span>
                    </label>
                  ))}
                </div>

                {/* Error Message Styling */}
                {field.state.meta.errors && (
                  <em className="text-xs font-medium text-red-500 mt-1 block">
                    {field.state.meta.errors.join(', ')}
                  </em>
                )}
              </div>
            )}
          />
        </div>

        {/* Form Submission */}
        <div className="pt-4">
          <form.AppForm>
            <form.SubscribeButton label="Save Changes" />
          </form.AppForm>
        </div>
      </form>
    </div>
  )
}
