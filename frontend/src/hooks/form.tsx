import { createFormHook } from '@tanstack/react-form'
import { lazy } from 'react'
import { fieldContext, formContext, useFormContext } from './form-context.tsx'

const TextField = lazy(() => import('../components/text-fields.tsx'))
const TextArea = lazy(() => import('../components/text-area.tsx'))


function SubscribeButton({ label }: { label: string }) {
  const form = useFormContext()
  return (
    <form.Subscribe selector={(state) => state.isSubmitting}>
      {(isSubmitting) => <button disabled={isSubmitting}>{label}</button>}
    </form.Subscribe>
  )
}

export const { useAppForm, withForm, withFieldGroup } = createFormHook({
  fieldComponents: {
    TextField,
    TextArea
  },
  formComponents: {
    SubscribeButton,
  },
  fieldContext,
  formContext,
})
