"use client";

import Image from "next/image";
// import CoffeeImage from "@/public/product-image/coffee.png";
import { formatNumber } from "@/lib/utils";
import { useCartStore } from "@/hooks/use-cart";
import { FoodCartProps } from "./show-items";

export default function FoodCard({ foods }: { foods: FoodCartProps }) {

  const { add: handleAddToCart } = useCartStore();

  return (
    <div className="border p-3 rounded-xl border-slate-700">
      <div className="bg-gray-300 rounded-md mb-2">
        {/* <img
          src="https://github.com/SAF2k/ecommerce-store/blob/main/public/billboard-bg-3.png"
          alt="coffee"
          className="w-[180px] h-[180px] rounded object-cover"
        /> */}
      </div>
      <h2 className="text-slate-400">{foods.name}</h2>
      <h2 className="font-semibold text-green-400">$ {formatNumber(foods.price)}</h2>
      <button
        onClick={() => handleAddToCart(foods)}
        className="mt-4 font-semibold text-sm bg-slate-100 text-slate-800 rounded-md py-2 text-center w-full"
      >
        Add To Cart
      </button>
    </div>
  );
}
