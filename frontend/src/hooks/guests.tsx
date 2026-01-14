import { useQuery, UseQueryResult, useMutation } from '@tanstack/react-query';

import {createGuest} from './identity';
import { Party, Guest } from '../interfaces/party';
import { Response } from '../interfaces/response';

export const getGuests = (code: string): UseQueryResult<Guest[]> => {
    return useQuery({
        queryKey: ['guests'],
        queryFn: async (): Promise<Guest[]> => {
            const response = await fetch(`${import.meta.env.VITE_API_URL}/guests/${code}`);
            return await response.json()
        }
    })
}

