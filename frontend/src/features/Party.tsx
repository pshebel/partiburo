import { useAppForm } from '../hooks/form.tsx'
import {  useMutation } from '@tanstack/react-query';
import { Party } from '../interfaces/party'
import { Response } from '../interfaces/response'
import { useNavigate } from 'react-router-dom'
import { formOptions } from '@tanstack/react-form'

const partyFormOptions = formOptions({
  defaultValues: {
    admin_email: '',
    title: '',
    description: '',
    location: '',
    date: '',
    time: '',
    reminders: [],
    announcements: [],
  },
  validators: {
    onChangeAsync: async ({ value }) => {
      const errors = {
        fields: {},
      } as {
        fields: Record<string, string>
      }
      // if (!value.title) {
      //   errors.fields.title = 'Title is required'
      // }
      // if (!value.description) {
      //   errors.fields.description = 'Description is required'
      // }
      // if (!value.location) {
      //   errors.fields.location = 'Location is required'
      // }
      // if (!value.date) {
      //   errors.fields.date = 'Date is required'
      // }


      if (value.title.length > 500) {
        errors.fields.name = 'Title is too long. Must be less than 500 characters.'
      }
      
      if (value.description.length > 5000) {
        errors.fields.description = 'Description is too long. Must be less than 5000 characters.'
      }

     

      return errors
    },
  },
})

export const CreateParty = () => {
  const navigate = useNavigate()

  const form = useAppForm({
    ...partyFormOptions,
    onSubmit: async ({ formApi, value }) => {
      await saveUserMutation.mutateAsync(value)

      // Reset the form to start-over with a clean state
      formApi.reset()
    },
  })

  const saveUserMutation = useMutation({
    mutationFn: async (req: { body: Party;}) => {
        const response = await fetch(`${import.meta.env.VITE_API_URL}/party`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                },
                body: JSON.stringify(req),
        });
        return response.json() as Promise<Response>;
    },
    onSuccess: (data) => {
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
        className="max-w-xl w-full bg-white p-8 rounded-2xl shadow-sm border border-gray-100 space-y-6"
      >
        <div>
          <form.AppField
            name="admin_email"
            children={(field) => (
              <field.TextField 
                label="Admin Email" 
              />
            )}
          />
        </div>

        <div>
          <form.AppField
            name="title"
            children={(field) => (
              <field.TextField 
                label="Title" 
              />
            )}
          />
        </div>

        <div>
          <form.AppField
            name="description"
            children={(field) => (
              <field.TextArea 
                label="Description" 
              />
            )}
          />
        </div>

        <div>
          <form.AppField
            name="location"
            children={(field) => (
              <field.TextField 
                label="Location" 
              />
            )}
          />
        </div>
        <div>
          <form.AppField
            name="date"
            children={(field) => (
              <field.DateField 
                label="Date" 
              />
            )}
          />
        </div>
        <div>
          <form.AppField
            name="time"
            children={(field) => (
              <field.TimeField 
                label="Time" 
              />
            )}
          />
        </div>

        <div>
          <form.AppField
            name="reminders"
            children={(field) => (
              <field.MultiSelectField
                label="reminders"
                options={[{
                        label: 'day of',
                        value: 'dayof',
                    },
                    {
                        label: 'day before',
                        value: 'daybefore',
                    },
                    {
                        label: '3 days before',
                        value: 'threedaysbefore',
                    },
                    {
                        label: 'week before',
                        value: 'weekbefore',
                    },
                    {
                      label: 'announcements',
                      value: 'announcements',
                    }
                ]}
              />
            )}
          />
        </div>
        {/* Form Submission */}
        <div className="pt-4">
          <form.AppForm>
            <form.SubscribeButton label="Create Party" />
          </form.AppForm>
        </div>
      </form>
    </div>
  )
}
