import { formOptions } from '@tanstack/react-form'

const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/

export const loginFormOptions = formOptions({
  defaultValues: {
    status: '',
    email: '',
  },
  validators: {
    // Synchronous validation is much better for "as-you-type" logic
    onChange: ({ value }) => {
      if (value.email && !emailRegex.test(value.email)) {
        return {
          fields: {
            email: 'Please enter a valid email address',
          },
        }
      }
      return undefined
    },
  },
})