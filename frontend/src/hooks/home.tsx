import { useQuery, UseQueryResult, useMutation } from '@tanstack/react-query';

import { Home } from '../interfaces/party';


export const getHome = (code: string): UseQueryResult<Home> => {
    return useQuery({
        queryKey: ['home', code],
        queryFn: async (): Promise<Home> => {
            const response = await fetch(`${import.meta.env.VITE_API_URL}/home/${code}`);
            if (!response.ok) {
                const errorData: Response = await response.json();
                throw { ...errorData, status: response.status };
            }
            return await response.json()
        }
    })
}


