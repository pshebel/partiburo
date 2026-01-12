import {useState, ChangeEvent } from 'react'
import { useAppForm } from '../../hooks/form.tsx'
import { loginFormOptions } from './login-form-options.tsx'
import {  useMutation } from '@tanstack/react-query';
import { getGuests } from '../../hooks/guests';
import { createGuest } from '../../hooks/identity';
import {Guest} from '../../interfaces/party';
import { useNavigate } from 'react-router-dom'

export const SelectGuest = () => {
  const navigate = useNavigate()

  const [guest, setGuest] = useState('');
  const { data, isLoading, error } = getGuests();

  if (isLoading) {
    return (<></>)
  }

  if (error) {
    return(<></>)
  }


  const handleSubmit = () => {
    console.log('handle submit ', guest)
    if (guest === '') {
      return
    }
    createGuest(guest)
    navigate('/')
  }

  const handleChange = (e: ChangeEvent<HTMLSelectElement>) => {
    setGuest(e.target.value)
  }

  console.log(data)
  return (
    <div className="bg-gray-100/50 p-8 rounded-2xl border border-dashed border-gray-300">
      <h1 className="text-lg font-bold text-gray-800 mb-4 text-center">Already RSVP'd?</h1>
      <div className="flex flex-col gap-3">
        <select 
          value={guest} 
          onChange={handleChange}
          className="w-full p-3 rounded-xl border border-gray-200 bg-white focus:ring-2 focus:ring-gray-400 outline-none transition"
        >
          <option value="">Select your name...</option>
          {data.map((g: Guest) => (
            <option key={g.id} value={g.id}>{g.name}</option>
          ))}
        </select>
        
        <button 
          onClick={handleSubmit}
          disabled={!guest}
          className="w-full bg-gray-800 text-white font-bold py-3 rounded-xl hover:bg-black transition disabled:opacity-50 disabled:cursor-not-allowed"
        >
          Continue as Guest
        </button>
      </div>
    </div>
  )
}
