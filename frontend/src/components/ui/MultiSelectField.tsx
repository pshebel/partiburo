import { useStore } from '@tanstack/react-form'
import { useFieldContext } from '../../hooks/form-context.tsx'

interface Option {
  label: string
  value: string
}

export default function MultiSelectField({ 
  label, 
  options 
}: { 
  label: string; 
  options: Option[] 
}) {
  const field = useFieldContext<string[]>()
  const errors = useStore(field.store, (state) => state.meta.errors)
  
  const selectedValues = field.state.value ?? []

  const toggleOption = (val: string) => {
    if (selectedValues.includes(val)) {
      field.handleChange(selectedValues.filter((v) => v !== val))
    } else {
      field.handleChange([...selectedValues, val])
    }
  }

  return (
    <div className="flex flex-col gap-1.5">
      <label className="text-sm font-semibold text-gray-700 mb-1">
        {label}
      </label>
      
      <div className="min-h-[3rem] w-full p-2 rounded-xl border border-gray-200 bg-white shadow-sm focus-within:ring-2 focus-within:ring-blue-500 focus-within:border-transparent transition-all flex flex-wrap gap-2">
        {/* Render Selected Tags */}
        {selectedValues.map((val) => (
          <span 
            key={val} 
            className="flex items-center gap-1.5 bg-blue-50 text-blue-700 px-2.5 py-1 rounded-lg text-sm font-medium border border-blue-100"
          >
            {options.find(o => o.value === val)?.label || val}
            <button
              type="button"
              onClick={() => toggleOption(val)}
              className="hover:bg-blue-200 text-blue-400 hover:text-blue-800 rounded-md transition-colors leading-none"
              aria-label="Remove"
            >
              {/* Inline SVG X icon */}
              <svg 
                viewBox="0 0 24 24" 
                width="14" 
                height="14" 
                stroke="currentColor" 
                strokeWidth="2.5" 
                fill="none" 
                strokeLinecap="round" 
                strokeLinejoin="round"
              >
                <line x1="18" y1="6" x2="6" y2="18"></line>
                <line x1="6" y1="6" x2="18" y2="18"></line>
              </svg>
            </button>
          </span>
        ))}

        {/* Selection Input */}
        <select
          className="bg-transparent outline-none text-sm text-gray-400 cursor-pointer flex-1 min-w-[120px] h-8"
          value=""
          onChange={(e) => {
            if (e.target.value) toggleOption(e.target.value)
          }}
          onBlur={field.handleBlur}
        >
          <option value="" disabled>Add item...</option>
          {options
            .filter(opt => !selectedValues.includes(opt.value))
            .map((opt) => (
              <option key={opt.value} value={opt.value}>
                {opt.label}
              </option>
            ))}
        </select>
      </div>

      {errors.map((error: any) => (
        <div key={error} className="text-xs font-medium text-red-500 mt-1 ml-1 animate-in fade-in slide-in-from-top-1">
          {error}
        </div>
      ))}
    </div>
  )
}