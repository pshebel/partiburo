import { useQuery, UseQueryResult, useMutation } from '@tanstack/react-query';

import { Party } from '../interfaces/party';


export const getParty = (): UseQueryResult<Party> => {
    return useQuery({
        queryKey: ['party'],
        queryFn: async (): Promise<Party> => {
            const response = await fetch(`${import.meta.env.VITE_API_URL}/party`);
            return await response.json()
        }
    })
}


