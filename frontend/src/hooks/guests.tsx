import { useQuery, UseQueryResult, useMutation } from '@tanstack/react-query';
import { Guest } from '../interfaces/party';
import { Response } from '../interfaces/response';

export const getGuests = (code: string): UseQueryResult<Guest[]> => {
    return useQuery({
        queryKey: ['guests', code],
        queryFn: async (): Promise<Guest[]> => {
            const response = await fetch(`${import.meta.env.VITE_API_URL}/guests/${code}`);
            if (!response.ok) {
                const errorData: Response = await response.json();
                throw { ...errorData, status: response.status };
            }
            return await response.json()
        }
    })
}

