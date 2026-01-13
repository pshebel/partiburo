import { formOptions } from '@tanstack/react-form'
const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/

export const loginFormOptions = formOptions({
  defaultValues: {
    name: '',
    email: '',
    status: '',
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
      return errors
    },
  },
})
