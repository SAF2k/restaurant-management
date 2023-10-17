"use client";

import { formatNumber } from "@/lib/utils";
import { useCartStore } from "@/hooks/use-cart";
import { FoodCartProps } from "./show-items";
import FoodsItemImage from "./food-image";
import { Button } from "@/components/ui/button";

export default function FoodCard({ foods }: { foods: FoodCartProps }) {

  const { add: handleAddToCart } = useCartStore();

  return (
    <div
      className="border border-slate-600 dark:border-slate-400 p-3 rounded-xl"
      onClick={() => handleAddToCart(foods)}
    >
      <div className="bg-current rounded-md mb-2">
        <FoodsItemImage />
      </div>
      <h2 className="text-base font-semibold">{foods.name}</h2>
      <h2 className="font-semibold text-green-400">
        ₹ {formatNumber(foods.price)}
      </h2>
      <Button className="mt-4 rounded-md py-2 text-center w-full">
        Add To Cart
      </Button>
    </div>
  );
}
