import { useAppForm } from '../../hooks/form.tsx'
import { postFormOptions } from './post-form-options.tsx'
import {  useMutation } from '@tanstack/react-query';
import { getGuest } from '../../hooks/identity'
import { Response } from '../../interfaces/response.js'
import { useParams, useNavigate } from 'react-router-dom'

export const Post = () => {
  const navigate = useNavigate()
  const { code } = useParams();
  if (code === undefined) {
      navigate('/')
  }

  const guest_id = getGuest(code);
  if (guest_id === null) {
      navigate(`/login/${code}`)
  }
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
        const response = await fetch(`${import.meta.env.VITE_API_URL}/post/${code}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                },
                body: JSON.stringify(r),
        });
        return response.json() as Promise<Response>;
    },
    onSuccess: () => {
      navigate(`/${code}`)
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
        className="max-w-xl w-full bg-white p-8 rounded-2xl shadow-sm border border-gray-100 space-y-6"
      >
        <div className="border-b pb-4 mb-2">
          <h1 className="text-xs font-bold uppercase tracking-widest text-purple-600 mb-1">Community</h1>
          <h2 className="text-2xl font-extrabold text-gray-900">Create Post</h2>
        </div>

        <div className="space-y-4">
          <form.AppField
            name="body"
            children={(field) => (
              <field.TextArea 
                label="What's on your mind?" 
              />
            )}
          />
        </div>

        <div className="pt-2">
          <form.AppForm>
            <form.SubscribeButton label="Post to Board" />
          </form.AppForm>
        </div>
      </form>
    </div>
  )
}
