import { getFoodById } from "@/actions/get-food";
import { FoodForm } from "./components/food-form";
import { getMenus } from "@/actions/get-menu";

const FoodPage = async ({
  params,
}: {
  params: { foodId: string; storeId: string };
}) => {
  const id = params.foodId.toString();
  const storeId = params.storeId.toString();

  // const foodData = async () => {
  //   if (id === "new") return null;
  //   console.log(id);

  //   const food = await getFoodById({ storeId });
  //   return food;
  // };
  const menus = await getMenus({ storeId });

  const food = params.foodId === "new" ? null : await getFoodById({ id,storeId });

  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <FoodForm menus={menus} initialData={food} />
      </div>
    </div>
  );
};

export default FoodPage;
