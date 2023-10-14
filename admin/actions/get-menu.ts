import axios from "axios";

export interface FoodData {
  _id: string;
  food_id: string;
  name: string;
  price: number;
  food_image: string;
}

export const getMenu = async (): Promise<FoodData[]> => {
  try {
    const response = await axios.get("http://localhost:8080/foods");
    const data = response.data; // Assuming response data is an array

    if (Array.isArray(data)) {
      return data;
    } else {
      console.error("Response data is not an array:", data);
      return [];
    }
  } catch (error) {
    // Handle any potential errors here
    console.error("Error fetching food data:", error);
    return []; // Return an empty array or handle the error as needed
  }
};
