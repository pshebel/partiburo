import { useStore } from '@tanstack/react-form'
import { useFieldContext } from '../../hooks/form-context.tsx'
import { Minus, Plus } from 'lucide-react'

interface NumberFieldProps {
  label: string
  step?: number
  min?: number
  max?: number
}

export default function NumberField({ label, step = 1, min, max }: NumberFieldProps) {
  const field = useFieldContext<number>()
  const errors = useStore(field.store, (state) => state.meta.errors)
  const currentValue = field.state.value ?? 0

  const adjustValue = (amount: number) => {
    const newValue = currentValue + amount
    
    // Clamp the value between min and max
    if (min !== undefined && newValue < min) return
    if (max !== undefined && newValue > max) return
    
    field.handleChange(newValue)
  }

  const handleInputChange = (val: number) => {
    let sanitizedValue = val
    
    // Logic for manual typing
    if (min !== undefined && sanitizedValue < min) sanitizedValue = min
    if (max !== undefined && sanitizedValue > max) sanitizedValue = max
    
    field.handleChange(sanitizedValue)
  }

  // Determine if buttons should be visually disabled
  const isMinDisabled = min !== undefined && currentValue <= min
  const isMaxDisabled = max !== undefined && currentValue >= max

  return (
    <div className="flex flex-col gap-1.5">
      <label className="text-sm font-semibold text-gray-700">
        {label}
      </label>
      
      <div className="flex items-center gap-2">
        <button
          type="button"
          disabled={isMinDisabled}
          onClick={() => adjustValue(-step)}
          className="p-3 rounded-xl border border-gray-200 bg-white hover:bg-gray-50 active:scale-95 transition-all shadow-sm text-gray-600 disabled:opacity-30 disabled:cursor-not-allowed disabled:active:scale-100"
        >
          <Minus size={18} />
        </button>

        <input
          type="number"
          min={min}
          max={max}
          className="w-full p-3 text-center rounded-xl border border-gray-200 bg-white focus:ring-2 focus:ring-blue-500 focus:border-transparent outline-none transition-all shadow-sm [appearance:textfield] [&::-webkit-outer-spin-button]:appearance-none [&::-webkit-inner-spin-button]:appearance-none"
          value={currentValue}
          onChange={(e) => handleInputChange(e.target.valueAsNumber || 0)}
          onBlur={field.handleBlur}
        />

        <button
          type="button"
          disabled={isMaxDisabled}
          onClick={() => adjustValue(step)}
          className="p-3 rounded-xl border border-gray-200 bg-white hover:bg-gray-50 active:scale-95 transition-all shadow-sm text-gray-600 disabled:opacity-30 disabled:cursor-not-allowed disabled:active:scale-100"
        >
          <Plus size={18} />
        </button>
      </div>

      {errors.map((error: string) => (
        <div key={error} className="text-xs font-medium text-red-500 mt-1 ml-1 animate-in fade-in slide-in-from-top-1">
          {error}
        </div>
      ))}
    </div>
  )
}