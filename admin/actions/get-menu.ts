import axios from "axios";

export interface MenuData {
  _id: string;
  name: string;
  category: string;
  created_at: string;
}

export const getMenus = async ({ storeId }: {storeId:string}) => {
  console.log("storeId", storeId);

  try {
    const response = await axios.get(
      `${process.env.NEXT_PUBLIC_API_URL}/${storeId}/menu`
    );
    const data = response.data;

    if (Array.isArray(data)) {
      const transformedData = data.map(
        ({ _id, name, category, created_at }) => ({
          _id,
          name,
          category,
          created_at,
        })
      );

      return transformedData;
    } else {
      console.error("Response data is not an array:", data);
      return [];
    }
  } catch (error) {
    console.error("Error fetching menu data:", error);
    return [];
  }
};

export const getMenuById = async ({ id, storeId }: {id:string, storeId:string}) => {
  try {
    const response = await axios.get(
      `${process.env.NEXT_PUBLIC_API_URL}/${storeId}/menu/${id}`
    );
    const data = response.data;

    return data;
  } catch (error) {
    console.error("Error fetching menu data:", error);
    return;
  }
};
