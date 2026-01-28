import { useQuery, UseQueryResult } from '@tanstack/react-query';

import { Party, TitleRequest, TitlesResponse } from '../interfaces/party';
import { Response }  from '../interfaces/response'


export const getTitles = (req: TitleRequest, isEnabled: boolean): UseQueryResult<TitlesResponse> => {
    return useQuery({
        queryKey: ['titles', req],
        queryFn: async (): Promise<TitlesResponse> =>{
            const response = await fetch(`${import.meta.env.VITE_API_URL}/titles`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(req),
            });

            if (!response.ok) {
                const errorData: Response = await response.json();
                throw { ...errorData, status: response.status };
            }
            return response.json();
        },
        enabled: isEnabled
    })
}

export const getParty = (code: string): UseQueryResult<Party> => {
    return useQuery({
        queryKey: ['party', code],
        queryFn: async (): Promise<Party> => {
            const response = await fetch(`${import.meta.env.VITE_API_URL}/party/${code}`);
            if (!response.ok) {
                const errorData: Response = await response.json();
                throw { ...errorData, status: response.status };
            }
            return await response.json()
        }
    })
}


