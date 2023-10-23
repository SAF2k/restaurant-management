import { create } from "zustand";
import { toast } from "react-hot-toast";
import { FoodCartProps } from "@/app/(dashboard)/[storeId]/(routes)/placeorder/components/show-items";

interface CartItem extends FoodCartProps {
  quantity: number;
}

export type CartStore = {
  cart: CartItem[];
  quantity: number;
  add: (food: FoodCartProps) => void;
  remove: (id: string) => void;
  removeAll: () => void;
};

export const useCartStore = create<CartStore>((set, get) => ({
  cart: [], // An array to store the items in the cart
  quantity: 0,
  add: (food: FoodCartProps) => {
    const { cart } = get();
    const existingItemIndex = cart.findIndex((item) => item.ID === food.ID);
    if (existingItemIndex !== -1) {
      // If the item already exists, increase its quantity
      const updatedCart = [...cart];
      updatedCart[existingItemIndex].quantity++;
      set({ cart: updatedCart });
    } else {
      // If it's a new item, add it to the cart with quantity 1
      set({ cart: [...cart, { ...food, quantity: 1 }] });
      toast.success("New Item added to cart");
    }
  },
  remove: (id: string) => {
    const { cart } = get();
    let itemRemoved = false;
    const updatedCart = cart
      .map((item) => {
        if (itemRemoved || item.ID !== id) {
          return item;
        }
        if (item.quantity > 1) {
          itemRemoved = true;
          return { ...item, quantity: item.quantity - 1 };
        }
        itemRemoved = true;
        toast.error("All " + item.name + " removed");
        return null;
      })
      .filter(Boolean) as CartItem[]; // Use type assertion here
    set({ cart: updatedCart });
  },
  removeAll: () => {
    toast.error("All items removed");
    set({ cart: [] });
  },
}));
