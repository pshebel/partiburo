import { useStore } from '@tanstack/react-form'
import { useFieldContext } from '../hooks/form-context.tsx'

export default function TextField({ label }: { label: string }) {
  const field = useFieldContext<string>()
  const errors = useStore(field.store, (state) => state.meta.errors)

  return (
    <div className="flex flex-col gap-1.5">
      <label className="group">
        <div className="text-sm font-semibold text-gray-700 mb-1 transition-colors group-focus-within:text-blue-600">
          {label}
        </div>
        <input
          className="w-full p-3 rounded-xl border border-gray-200 bg-white focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none transition-all placeholder:text-gray-400 shadow-sm"
          value={field.state.value}
          onChange={(e) => field.handleChange(e.target.value)}
          onBlur={field.handleBlur}
        />
      </label>
      {errors.map((error: string) => (
        <div key={error} className="text-xs font-medium text-red-500 mt-1 ml-1 animate-in fade-in slide-in-from-top-1">
          {error}
        </div>
      ))}
    </div>
  )
}