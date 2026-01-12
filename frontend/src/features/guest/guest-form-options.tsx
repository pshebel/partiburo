import { formOptions } from '@tanstack/react-form'

export const loginFormOptions = formOptions({
  defaultValues: {
    status: '',
    email: '',
  },
//   validators: {
//     onChangeAsync: async ({ value }) => {
//       const errors = {
//         fields: {},
//       } as {
//         fields: Record<string, string>
//       }
//       if (!value.name ) {
//         errors.fields.name = 'Name is required'
//       }
//       if (!value.status) {
//         errors.fields.name = 'Status is required'
//       }
//       return errors
//     },
//   },
})
