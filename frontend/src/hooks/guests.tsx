import { useQuery, UseQueryResult, useMutation } from '@tanstack/react-query';

import {createGuest} from './identity';
import { Party, Guest } from '../interfaces/party';
import { Response } from '../interfaces/response';

export const getGuests = (): UseQueryResult<Guest[]> => {
    return useQuery({
        queryKey: ['guests'],
        queryFn: async (): Promise<Guest[]> => {
            // const response = await fetch(`${process.env.API_URL}`);
            const response = await fetch('http://localhost:4000/guests');
            return await response.json()
        }
    })
}
