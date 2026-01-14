import { useQuery, UseQueryResult } from '@tanstack/react-query';

import { Party, TitleRequest, TitlesResponse } from '../interfaces/party';


export const getTitles = (req: TitleRequest): UseQueryResult<TitlesResponse> => {
    return useQuery({
        queryKey: ['titles'],
        queryFn: async (): Promise<TitlesResponse> =>{
            const response = await fetch(`${import.meta.env.VITE_API_URL}/titles`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(req),
            });
            return response.json();
        }
    })
}

export const getParty = (code: string): UseQueryResult<Party> => {
    return useQuery({
        queryKey: ['party'],
        queryFn: async (): Promise<Party> => {
            const response = await fetch(`${import.meta.env.VITE_API_URL}/party/${code}`);
            return await response.json()
        }
    })
}


