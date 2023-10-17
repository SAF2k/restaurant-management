"use client";
import { useCartStore } from "@/hooks/use-cart";

export default function CartItems() {
  const cart = useCartStore();

  console.log(cart.cart);

  return (
    <div className="flex flex-col gap-2 mt-4">
      <table className="text-slate-400">
        <thead>
          <tr>
            <th>#</th>
            <th>Item</th>
            <th>Price</th>
            <th>Count</th>
            <th>Total</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td colSpan={5} className="text-center text-slate-600 py-2">
              No Data
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  );
}
