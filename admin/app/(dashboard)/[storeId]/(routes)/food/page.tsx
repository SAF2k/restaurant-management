import { FoodClient } from "./components/client";


const FoodsPage = () => {

  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <FoodClient />
      </div>
    </div>
  );
};

export default FoodsPage;
