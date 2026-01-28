import { useStore } from '@tanstack/react-form'
import { useFieldContext } from '../../hooks/form-context.tsx'

export default function DateField({ label }: { label: string }) {
  // TanStack Form usually stores dates as ISO strings (YYYY-MM-DD) for native inputs
  const field = useFieldContext<string>()
  const { errors, isTouched } = useStore(field.store, (state) => ({
    errors: state.meta.errors,
    isTouched: state.meta.isTouched,
  }))

  const shouldShowError = isTouched && errors.length > 0

  return (
    <div className="flex flex-col gap-1.5">
      <label className="group">
        {/* Dynamic label color */}
        <div className={`text-sm font-semibold mb-1 transition-colors ${
          shouldShowError ? 'text-red-500' : 'text-gray-700 group-focus-within:text-blue-600'
        }`}>
          {label}
        </div>
        
        <input
          type="date"
          className={`w-full p-3 rounded-xl border bg-white outline-none transition-all placeholder:text-gray-400 shadow-sm ${
            shouldShowError 
              ? 'border-red-500 focus:ring-2 focus:ring-red-200' 
              : 'border-gray-200 focus:ring-2 focus:ring-blue-500 focus:border-transparent'
          }`}
          // Native date inputs expect YYYY-MM-DD format
          value={field.state.value ?? ''}
          onChange={(e) => field.handleChange(e.target.value)}
          onBlur={field.handleBlur}
        />
      </label>
      {/* 3. Conditional Error Message Rendering */}
      {shouldShowError && errors.map((error: any) => (
        <div 
          key={typeof error === 'string' ? error : error.message} 
          className="text-xs font-medium text-red-500 mt-1 ml-1 animate-in fade-in slide-in-from-top-1"
        >
          {typeof error === 'string' ? error : error.message}
        </div>
      ))}
    </div>
  )
}