import { useQuery, UseQueryResult, useMutation } from '@tanstack/react-query';
import { UnsubscribeRequest } from '../interfaces/unsubscribe';
import { Response } from '../interfaces/response';


export const postUnsubscribe = (party_code: string, email_code: string): UseQueryResult<Response> => {
    const body: UnsubscribeRequest = {
        party_code: party_code,
        email_code: email_code,
        all: false,
    }
    return useQuery({
        queryKey: ['unsubscribe', body],
        queryFn: async (): Promise<Response> => {
            const response = await fetch(`${import.meta.env.VITE_API_URL}/unsubscribe`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(body),
            });
            if (!response.ok) {
                const errorData: Response = await response.json();
                throw { ...errorData, status: response.status };
            }
            return await response.json() as Promise<Response>;
        }
    })
}


export const postUnsubscribeAll = (email_code: string): UseQueryResult<Response> => {
    const body: UnsubscribeRequest = {
        party_code: '',
        email_code: email_code,
        all: true,
    }

    
    return useQuery({
        queryKey: ['unsubscribe', email_code],
        queryFn: async (): Promise<Response> => {
            const response = await fetch(`${import.meta.env.VITE_API_URL}/unsubscribe`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(body),
            });
            if (!response.ok) {
                const errorData: Response = await response.json();
                throw { ...errorData, status: response.status };
            }
            return await response.json() as Promise<Response>;
        }
    })
}
