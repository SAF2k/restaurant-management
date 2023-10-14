"use client"

import { Tabs, TabsContent } from "@radix-ui/react-tabs";
import FoodItem from "./food-items";
import { useEffect, useState } from "react";
import { FoodData, getFood } from "@/actions/get-food";

export default function FoodCollection() {

    const [foodData, setFoodData] = useState<FoodData[]>([]);

    useEffect(() => {
        const getFoodData = async () => {
            const foodData = await getFood();
            setFoodData(foodData);
        };
        getFoodData();
    }, []);
    return (
      <div className="grid gap-4 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-5 p-2 md:p-4 lg:p-10">
        {foodData.map((food) => (
          <Tabs
            defaultValue="overview"
            className="space-y-4"
            key={food.food_id}
          >
            <TabsContent value="overview" className="space-y-4">
              <div className="">
                <FoodItem key={food.food_id} food={food} />
              </div>
            </TabsContent>
          </Tabs>
        ))}
      </div>
    );
}