import { useQuery, UseQueryResult, useMutation } from '@tanstack/react-query';

import { Home } from '../interfaces/party';


export const getHome = (code: string): UseQueryResult<Home> => {
    return useQuery({
        queryKey: ['home'],
        queryFn: async (): Promise<Home> => {
            const response = await fetch(`${import.meta.env.VITE_API_URL}/home/${code}`);
            return await response.json()
        }
    })
}


