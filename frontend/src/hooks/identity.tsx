
export const getGuest = () => {
    return localStorage.getItem('guest_id') || null;
}

export const createGuest = (id: string) => {
    return localStorage.setItem('guest_id', id);
}
