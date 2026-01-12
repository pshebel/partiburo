import { useAppForm } from '../../hooks/form.tsx'
import { loginFormOptions } from './login-form-options.tsx'
import {  useMutation } from '@tanstack/react-query';
import { useNavigate } from 'react-router-dom'
import { createGuest } from '../../hooks/identity';
import { GuestResponse } from '../../interfaces/response.js'

export const CreateGuest = () => {
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
    mutationFn: async (req: { name: string, email: string, status: string }) => {
        const response = await fetch(`${import.meta.env.VITE_API_URL}/guest`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                },
                body: JSON.stringify(req),
        });
        return response.json() as Promise<Response>;
    },
    onSuccess: (data: GuestResponse) => {
      createGuest(data.id)
      navigate('/')
    },
    onError: (err: any) => {
      window.confirm(err)
    },
  })

  return (
    <form
      onSubmit={(e) => { e.preventDefault(); form.handleSubmit(); }}
      className="bg-white p-8 rounded-2xl shadow-sm border border-gray-100"
    >
      <h1 className="text-xl font-bold text-gray-900 mb-6">RSVP / Identify Yourself</h1>
      
      <div className="space-y-5">
        <form.AppField
          name="name"
          children={(field) => (
            <div className="flex flex-col gap-1">
               <label className="text-sm font-semibold text-gray-700">Name</label>
               <field.TextField className="w-full p-3 rounded-xl border border-gray-200 focus:ring-2 focus:ring-blue-500 outline-none transition" />
            </div>
          )}
        />

        <form.AppField
          name="email"
          children={(field) => (
            <div className="flex flex-col gap-1">
               <label className="text-sm font-semibold text-gray-700">Email (Optional)</label>
               <field.TextField className="w-full p-3 rounded-xl border border-gray-200 focus:ring-2 focus:ring-blue-500 outline-none transition" placeholder="To receive updates" />
            </div>
          )}
        />

        <form.Field
          name="status"
          children={(field) => (
            <div className="space-y-2">
              <label className="text-sm font-semibold text-gray-700">Will you be attending?</label>
              <div className="grid grid-cols-3 gap-2">
                {['GOING', 'MAYBE', 'NOT_GOING'].map((val) => (
                  <label key={val} className={`
                    flex items-center justify-center p-3 rounded-xl border cursor-pointer transition text-sm font-bold
                    ${field.state.value === val ? 'bg-blue-600 border-blue-600 text-white' : 'bg-white border-gray-200 text-gray-600 hover:bg-gray-50'}
                  `}>
                    <input
                      type="radio"
                      className="hidden"
                      name={field.name}
                      value={val}
                      checked={field.state.value === val}
                      onChange={(e) => field.handleChange(e.target.value)}
                    />
                    {val.replace('_', ' ')}
                  </label>
                ))}
              </div>
              {field.state.meta.errors && <em className="text-xs text-red-500">{field.state.meta.errors.join(', ')}</em>}
            </div>
          )}
        />

        <div className="pt-4">
          <form.AppForm>
            <form.SubscribeButton label={"Submit"} className="w-full bg-blue-600 text-white font-bold py-3 rounded-xl hover:bg-blue-700 transition-all shadow-lg shadow-blue-100" />
          </form.AppForm>
        </div>
      </div>
    </form>
  )
}
