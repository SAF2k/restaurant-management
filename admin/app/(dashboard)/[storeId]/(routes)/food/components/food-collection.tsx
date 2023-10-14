"use client"

import { Tabs, TabsContent } from "@radix-ui/react-tabs";
import FoodItem from "./food-items";
import { useEffect, useState } from "react";
import { FoodData, getAllFood, getFoodByMenu,  } from "@/actions/get-food";

export default function FoodCollection({id} : {id?: string}) {

    const [foodData, setFoodData] = useState<FoodData[]>([]);

    useEffect(() => {
        const getFoodData = async () => {
          if (id) {
            const foodData = await getFoodByMenu({id});
            setFoodData(foodData);
            return;
          }else{
            const foodData = await getAllFood();
            setFoodData(foodData);
            return;
          }
        };
        getFoodData();
    }, [id]);
    
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