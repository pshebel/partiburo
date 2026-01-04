import { CreateGuest } from './create-guest';
import { SelectGuest } from './select-guest';


interface LoginProps {
    onLoginSuccess: (id: string) => void;
}

export const Login = ({ onLoginSuccess }: LoginProps) => {
  return (
    <div>
      <CreateGuest onLoginSuccess={onLoginSuccess}/>
      <SelectGuest onLoginSuccess={onLoginSuccess}/>
    </div>
  )
}
