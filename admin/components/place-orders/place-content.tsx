"use client"

import { useEffect, useState } from "react";
import { MenuData, getMenu } from "@/actions/get-menu";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "../ui/tabs";

export default function PlaceOrder() {

    const [menuData, setMenuData] = useState<MenuData[]>([]);

    useEffect(() => {
      const fetchMenuData = async () => {
        const menuItems: MenuData[] = await getMenu();
        setMenuData(menuItems);
      };
      fetchMenuData();
    }, []);

    return (
      <>
        <div className="flex mt-8">
          <Tabs defaultValue="special" className="flex flex-row">
            <TabsList className="flex flex-col gap-2 sm:w-40 lg:w-60 h-fit py-4 mr-5">
              {menuData.map((item) => (
                <TabsTrigger
                  value={item.category}
                  key={item._id}
                  className="w-full h-12 text-md"
                >
                  <>{item.category}</>
                </TabsTrigger>
              ))}
            </TabsList>
            {menuData.map((item) => (
              <TabsContent value={item.category} key={item._id}>
                {/* <FoodCollection id={item._id} /> */}
              </TabsContent>
            ))}
          </Tabs>
        </div>
      </>
    );
}