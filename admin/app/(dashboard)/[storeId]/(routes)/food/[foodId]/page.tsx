import { getFoodById } from "@/actions/get-food";
import { FoodForm } from "./components/food-form";
import { getMenus } from "@/actions/get-menu";

const FoodPage = async ({ params }: { params: { foodId: string } }) => {
  const foodData = async () => {
    if (params.foodId === "new") return null;
    const food = await getFoodById(params.foodId);
    return food;
  };
  const menus = await getMenus();

  const food = await foodData();

  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <FoodForm menus={menus} initialData={food} />
      </div>
    </div>
  );
};

export default FoodPage;
