import { ReactNode } from 'react'
import { Navigate } from 'react-router-dom'
import { getGuest } from '../hooks/identity'

interface ProtectedRouteProps {
  children: ReactNode
}

export const ProtectedRoute = ({ children }: ProtectedRouteProps) => {
  const navigate = useNavigate()
  const { code } = useParams();
  if (code === undefined) {
      navigate('/')
  }
  const guest_id = getGuest()
  
  if (guest_id === null || guest_id === "")  {
    return <Navigate to="/login" replace />
  }
  
  return <>{children}</>
}