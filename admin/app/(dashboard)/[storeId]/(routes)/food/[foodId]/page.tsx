"use client";

import axios from "axios";
import { FoodForm } from "./components/food-form";
import { useEffect, useState } from "react";
import { FoodData, getFoodById } from "@/actions/get-food";
import { MenuData } from "@/actions/get-menu";
import toast from "react-hot-toast";

const FoodPage = ({ params }: { params: { foodId: string } }) => {
  const [food, setFood] = useState<FoodData | null>(null);
  const [menus, setMenus] = useState<MenuData[]>([]);

  useEffect(() => {
    const fetchFood = async () => {
      try {
        if (params.foodId == "new") {
          setFood(null);
        } else {
          const foodData: any = await getFoodById(params.foodId);
          setFood(foodData);
        }
      } catch (error: any) {
        toast.error("Something went wrong.");
      }
    };
    const fetchMenus = async () => {
      try {
        const { data } = await axios.get(`http://localhost:8080/menus`);
        setMenus(data);
      } catch (error: any) {
        toast.error("Something went wrong.");
      }
    };
    fetchMenus();

    fetchFood();
  }, [params.foodId]);

  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <FoodForm menus={menus} initialData={food} />
      </div>
    </div>
  );
};

export default FoodPage;
