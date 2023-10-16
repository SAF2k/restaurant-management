import axios from "axios";

export interface TableData {
    _id: string;
    number_of_guests: number;
    table_number: number;
    created_at: string;
}

export const getAllTable = async () => {
  try {
    const response = await axios.get("http://localhost:8080/tables");
    const data = response.data;

    if (Array.isArray(data)) {
      const transformedData = data.map(
        ({ _id, number_of_guests, table_number, created_at }) => ({
          _id,
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

export const getTableById = async (id: string) => {
  try {
    const response = await axios.get(`http://localhost:8080/table/${id}`);
    const data = response.data;

    return data;
  } catch (error) {
    console.error("Error fetching table data:", error);
    return [];
  }
}