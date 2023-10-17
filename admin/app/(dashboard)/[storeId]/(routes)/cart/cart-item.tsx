import { useCartStore } from "@/hooks/use-cart";
import { FoodData } from "@/actions/get-food";

export default function CartItem({ food }: FoodData) {
  // Recover the store action to remove items from the cart
  const removeFromCart = useCartStore((state) => state.remove);

  return (
    <li className="flex justify-between items-center gap-4  mb-2 shadow-md p-4">
      <div className="flex items-center justify-around gap-4">
        <div>
          <h3 className="font-semibold">{food.name}</h3>
          <span className="text-sm text-gray-400">${food.price}</span>
          <h1 className="font-bold">{food.quantity}</h1>
        </div>
      </div>
      <div>
        <button
          title="Remove Item"
          className="text-red-500 hover:text-red-600 ml-4"
          onClick={() => removeFromCart(food)}
        ></button>
      </div>
    </li>
  );
}
