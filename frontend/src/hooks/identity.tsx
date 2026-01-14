// Define the structure of your storage object
type GuestMap = Record<string, string>;

const STORAGE_KEY = 'partiburo';

export const getCodes = (): string[] => {
  const rawData = localStorage.getItem(STORAGE_KEY);
  if (!rawData) return [];

  try {
    const data: GuestMap = JSON.parse(rawData);
    // Object.keys returns an array of the dictionary keys (the codes)
    return Object.keys(data);
  } catch (error) {
    console.error("Failed to parse guest data for keys", error);
    return [];
  }
};

/**
 * Retrieves a specific guest ID by its code.
 */
export const getGuest = (code: string): string | null => {
  const rawData = localStorage.getItem(STORAGE_KEY);
  if (!rawData) return null;

  try {
    const data: GuestMap = JSON.parse(rawData);
    return data[code] || null;
  } catch (error) {
    console.error("Failed to parse guest data from localStorage", error);
    return null;
  }
};


/**
 * Adds or updates a guest ID for a specific code.
 */
export const createGuest = (code: string, id: string): void => {
  const rawData = localStorage.getItem(STORAGE_KEY);
  let data: GuestMap = {};

  if (rawData) {
    try {
      data = JSON.parse(rawData);
    } catch (error) {
      console.error("Corrupted storage found. Resetting guest data.");
      data = {};
    }
  }

  // Update or add the new code/id pair
  data[code] = id;
  
  localStorage.setItem(STORAGE_KEY, JSON.stringify(data));
};