import { FoodData, getFoodByMenu } from "@/actions/get-food";
import FoodCard from "./item-card";

export interface FoodCartProps {
  ID: string;
  name: string;
  price: number;
}

export default async function ShowTest({ id }: { id: string }) {
  const foodData: FoodData[] = await getFoodByMenu({ id });

  const food: FoodCartProps[] = foodData.map((item) => ({
    ID: item.ID,
    name: item.name,
    price: item.price,
  }));

  console.log("show item food", food);
  
  return (
    <>
      <h1 className="font-semibold text-2xl border-b pb-4 border-b-slate-700">
        Foods
      </h1>

      <div className="text-sm pt-4 max-w-[70vw] flex gap-4 flex-wrap">
        {food.map((food, index) => (
          <div key={index}>
            <FoodCard foods={food} />
          </div>
        ))}
      </div>
    </>
  );
}
