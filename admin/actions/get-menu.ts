import axios from "axios";

export interface MenuData {
  _id: string;
  name: string;
  category: string;
}

export const getMenu = async () => {
  try {
    const response = await axios.get("http://localhost:8080/menus");
    const data = response.data;

    if (Array.isArray(data)) {
      const transformedData = data.map(({_id, name, category }) => ({
        _id,
        name,
        category,
      }));
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