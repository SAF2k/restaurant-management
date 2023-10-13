import { getFood } from "@/actions/get-food";
import FoodItem from "./components/food-items";
import { Tabs, TabsContent } from "@/components/ui/tabs";

const MenuPage = async () => {
  const foodData = await getFood();
  return (
    <div className="grid gap-4 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-5 p-2 md:p-4 lg:p-10">
      {foodData.map((food) => (
        <Tabs defaultValue="overview" className="space-y-4" key={food.food_id}>
          <TabsContent value="overview" className="space-y-4">
            <div className="">
              <FoodItem key={food.food_id} food={food} />
            </div>
          </TabsContent>
        </Tabs>
      ))}
    </div>
  );
};

export default MenuPage;
