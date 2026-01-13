import { createFormHook } from '@tanstack/react-form'
import { lazy } from 'react'
import { fieldContext, formContext, useFormContext } from './form-context.tsx'

const TextField = lazy(() => import('../components/ui/TextField.tsx'))
const TextArea = lazy(() => import('../components/ui/TextArea.tsx'))


function SubscribeButton({ label }: { label: string }) {
  const form = useFormContext()
  return (
    <form.Subscribe selector={(state) => [state.isSubmitting, state.canSubmit]}>
      {([isSubmitting, canSubmit]) => (
        <button
          type="submit"
          disabled={isSubmitting || !canSubmit}
          className={`
            w-full py-4 px-6 rounded-xl font-bold text-white transition-all
            flex items-center justify-center gap-2
            ${isSubmitting || !canSubmit 
              ? 'bg-gray-400 cursor-not-allowed' 
              : 'bg-blue-600 hover:bg-blue-700 active:scale-[0.98] shadow-lg shadow-blue-100 cursor-pointer'
            }
          `}
        >
          {isSubmitting ? (
            <>
              <div className="w-5 h-5 border-2 border-white/30 border-t-white rounded-full animate-spin" />
              <span>Processing...</span>
            </>
          ) : (
            label
          )}
        </button>
      )}
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
