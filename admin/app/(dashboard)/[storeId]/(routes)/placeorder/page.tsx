

import { MenuData, getMenus } from "@/actions/get-menu";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import CartItems from "./components/cart-items";
import ShowTest from "./components/show-items";
import { useEffect, useState } from "react";

const PlaceOrderPage = async () => {
  const menuData: MenuData[] = await getMenus();

  return (
    <div className="flex justify-between lg:px-16">
      <div className="p-8 pt-6">
        <Tabs defaultValue={menuData[0]._id}>
          <TabsList>
            {menuData.map((item) => (
              <TabsTrigger value={item._id} key={item._id}>
                <>{item.category}</>
              </TabsTrigger>
            ))}
          </TabsList>
          {menuData.map((item) => (
            <TabsContent value={item._id} key={item._id}>
              <ShowTest id={item._id} />
            </TabsContent>
          ))}
        </Tabs>
      </div>
      <CartItems />
    </div>
  );
};

export default PlaceOrderPage;
