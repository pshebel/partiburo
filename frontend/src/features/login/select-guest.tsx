import {useState, ChangeEvent } from 'react'
import { useAppForm } from '../../hooks/form.tsx'
import { loginFormOptions } from './login-form-options.tsx'
import {  useMutation } from '@tanstack/react-query';
import { getGuests } from '../../hooks/guests';
import { createGuest } from '../../hooks/identity';
import {Guest} from '../../interfaces/party';
import { Response } from '../../interfaces/response.js'


interface SelectGuestProps {
    onLoginSuccess: (id: string) => void;
}

export const SelectGuest = ({ onLoginSuccess }: SelectGuestProps) => {
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
    onLoginSuccess(guest)
  }

  const handleChange = (e: ChangeEvent<HTMLSelectElement>) => {
    setGuest(e.target.value)
  }

  console.log(data)
  return (
    <div>
      <h1>Already RSVP'd?</h1>
      <div>
        <label>
          
          <select value={guest} onChange={handleChange}>
            <option value="">-----</option>
            {data.map((guest: Guest) => {
              return (
                <option value={guest.ID}>{guest.Name}</option>
              )
            })}
          </select>
        </label>
      </div>
      <div>
        <button onClick={handleSubmit}>Select</button>
      </div>
    </div>
  )
}
