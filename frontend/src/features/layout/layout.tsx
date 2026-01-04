import { useState, useEffect } from 'react';
import { Login } from '../login/login';
import { Party } from '../party/party';
import { getGuest } from '../../hooks/identity';

export const Layout = () => {
    const [id, setId] = useState(() => getGuest());

    // This function will trigger the re-render when called
    const handleLoginSuccess = (newId: string) => {
        setId(newId);
    };

    if (id === null) {
        return (
            // Pass the handler as a prop
            <Login onLoginSuccess={handleLoginSuccess} />
        )
    }

    return <Party />
}