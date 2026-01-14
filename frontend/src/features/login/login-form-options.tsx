import { formOptions } from '@tanstack/react-form'
const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/

export const loginFormOptions = formOptions({
  defaultValues: {
    name: '',
    email: '',
    status: '',
    plus: 0,
  },
  validators: {
    onChangeAsync: async ({ value }) => {
      const errors = {
        fields: {},
      } as {
        fields: Record<string, string>
      }
      if (value.email && !emailRegex.test(value.email)) {
        return {
          fields: {
            email: 'Please enter a valid email address',
          },
        }
      }
      if (!value.name ) {
        errors.fields.name = 'Name is required'
      }
      if (!value.status) {
        errors.fields.name = 'Status is required'
      }
      if (value.plus < 0 || value.plus > 250) {
        return {
          fields: {
            plus: 'Plus one must be positive'
          }
        }
      }
      return errors
    },
  },
})
