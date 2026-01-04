import { useQuery, UseQueryResult, useMutation } from '@tanstack/react-query';

import {createGuest} from './identity';
import { Party, Guest } from '../interfaces/party';
import { Response } from '../interfaces/response';


export const getParty = (): UseQueryResult<Party> => {
    return useQuery({
        queryKey: ['party'],
        queryFn: async (): Promise<Party> => {
            // const response = await fetch(`${process.env.API_URL}`);
            const response = await fetch('http://localhost:4000');
            return await response.json()
        }
    })
}


