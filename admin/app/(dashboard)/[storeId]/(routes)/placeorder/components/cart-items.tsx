"use client";

import { ScrollArea } from "@/components/ui/scroll-area";
import { Separator } from "@/components/ui/separator";
import { useCartStore } from "@/hooks/use-cart";
import { DeleteAllButton, DeleteButton } from "./delete-button";
import ConfirmButton from "./confirm-button";
import { useEffect, useState } from "react";
import { TableData, getAllTable } from "@/actions/get-table";

export default function CartItems() {
  const [tables, setTables] = useState<TableData[]>([]);
  const [activate, setActivate] = useState<boolean>(true);
  const cartData = useCartStore();
  const cart = cartData.cart;

  const total = cart.reduce((acc, item) => {
    return acc + item.price * item.quantity;
  }, 0);

  useEffect(() => {
    if (total === 0) {
      setActivate(true);
    } else {
      setActivate(false);
    }
  }, [total]);

  useEffect(() => {
    const tableData = async () => {
      const res = await getAllTable();
      setTables(res);
    };
    tableData();
  }, []);

  return (
    <>
      <div className="flex flex-col justify-between border rounded-md h-[80vh] mt-14">
        <div>
          <div className="flex justify-between p-4">
            <h2 className="font-semibold text-3xl text-center">Cart Items</h2>
            <DeleteAllButton />
          </div>
          <Separator />
        </div>
        {cart.length >= 0 ? (
          <ScrollArea className="flex-1 w-[350px] py-4 px-8">
            {cart.map((item) => (
              <div key={item.ID}>
                <div className="flex justify-between items-center my-4">
                  <h2 className="">{item.name}</h2>
                  <div className="flex items-center space-x-8">
                    <h2 className="pr-4">{item.quantity}</h2>
                    <DeleteButton ID={item.ID} />
                  </div>
                </div>
                <Separator />
              </div>
            ))}
          </ScrollArea>
        ) : (
          <ScrollArea className="flex-1 w-[350px] py-4 px-8">
            <h2 className="flex-1 text-red-500">No Item Added</h2>
          </ScrollArea>
        )}
        <div>
          <Separator />
          <div className="flex items-center p-5 justify-between">
            <div className="flex gap-5 text-lg">
              <h2 className="font-semibold">Total</h2>
              <h2 className="font-bold">₹ {total}</h2>
            </div>
            <ConfirmButton activate={activate} tables={tables} />
          </div>
        </div>
      </div>
    </>
  );
}
