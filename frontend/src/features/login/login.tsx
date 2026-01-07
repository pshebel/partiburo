import { CreateGuest } from './create-guest';
import { SelectGuest } from './select-guest';

import { getParty } from '../../hooks/party';

// interface LoginProps {
//     onLoginSuccess: (id: string) => void;
// }{ onLoginSuccess }: LoginProps

export const Login = () => {
  const { data, isLoading, error } = getParty();
  if (isLoading) {
    return(
      <div>
        Loading...
      </div>
    )
  }
  if (error) {
    return (
      <div>
        Error {error.message}
      </div>
    )
  }
  if (data === undefined) {
    return (
      <div>
        Failed to get data
      </div>
    )
  }
  return (
    <div>
      <div>
          <h1>about</h1>
          <h2>{data.Title}</h2>
          <div>{data.Description}</div>
          <div>Date: {data.Date}</div>
          <div>Time: {data.Time}</div>
          <div>Address: {data.Address}</div>
      </div>
      <CreateGuest />
      <SelectGuest />
    </div>
  )
}
