import { useQuery, UseQueryResult } from '@tanstack/react-query';
import { ConfirmRequest } from '../interfaces/confirm';
import { Response } from '../interfaces/response';


export const createConfirm = (code: string, passcode: string): UseQueryResult<Response> => {
    const body: ConfirmRequest = {
        code: code,
        passcode: passcode
    }
    return useQuery({
        queryKey: ['confirm'],
        queryFn: async (): Promise<Response> => {
            const response = await fetch(`${import.meta.env.VITE_API_URL}/confirm`, {
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
