import { useQuery, UseQueryResult, useMutation } from '@tanstack/react-query';
import { UnsubscribeRequest } from '../interfaces/unsubscribe';
import { Response } from '../interfaces/response';


export const postUnsubscribe = (email: string): UseQueryResult<Response> => {
    const body: UnsubscribeRequest = {
        party_id: 0,
        email: email,
        all: false,
    }
    return useQuery({
        queryKey: ['unsubscribe'],
        queryFn: async (): Promise<Response> => {
            const response = await fetch(`${import.meta.env.VITE_API_URL}/unsubscribe`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(body),
            });
            return await response.json() as Promise<Response>;
        }
    })
}


export const postUnsubscribeAll = (email: string): UseQueryResult<Response> => {
    const body: UnsubscribeRequest = {
        party_id: 0,
        email: email,
        all: true,
    }

    
    return useQuery({
        queryKey: ['unsubscribe'],
        queryFn: async (): Promise<Response> => {
            const response = await fetch(`${import.meta.env.VITE_API_URL}/unsubscribe`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(body),
            });
            return await response.json() as Promise<Response>;
        }
    })
}
