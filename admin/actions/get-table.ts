import axios from "axios";

export interface TableData {
    table_id: string;
    number_of_guests: number;
    table_number: number;
    created_at: string;
}

export const getAllTable = async ({storeId}: {storeId: string}) => {
  try {
    const response = await axios.get(
      `${process.env.NEXT_PUBLIC_API_URL}/${storeId}/table`
    );
    const data = response.data;
    if (Array.isArray(data)) {
      const transformedData = data.map(
        ({ table_id, number_of_guests, table_number, created_at }) => ({
          table_id,
          table_number,
          number_of_guests,
          created_at,
        })
      );


      return transformedData;
    } else {
      console.error("Response data is not an array:", data);
      return [];
    }
  } catch (error) {
    console.error("Error fetching table data:", error);
    return [];
  }
};

export const getTableById = async ({storeId,id}: {storeId:string, id: string}) => {
  try {
    const response = await axios.get(
      `${process.env.NEXT_PUBLIC_API_URL}/${storeId}/table/${id}`
    );
    const data = response.data;

    return data;
  } catch (error) {
    console.error("Error fetching table data:", error);
    return [];
  }
}