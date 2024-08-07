import { FoodData, getFoodByMenu } from "@/actions/get-food";
import FoodCard from "./food-card";
import { Separator } from "@/components/ui/separator";
export interface FoodCartProps {
  ID: string;
  name: string;
  price: number;
}

export default async function ShowItems({ id, storeId }: { id: string, storeId: string }) {
  const foodData: FoodData[] = await getFoodByMenu({ id, storeId });

  const food: FoodCartProps[] = foodData.map((item) => ({
    ID: item.ID,
    name: item.name,
    price: item.price,
  }));

  console.log("show item food", food);

  return (
    <>
      <h1 className="font-semibold text-2xl pb-4">
        Foods
      </h1>

      <Separator />

      <div className="text-sm pt-4 lg:max-w-[70vw] flex gap-4 flex-wrap">
        {food.map((food, index) => (
          <div key={index}>
            <FoodCard foods={food} />
          </div>
        ))}
      </div>
    </>
  );
}
