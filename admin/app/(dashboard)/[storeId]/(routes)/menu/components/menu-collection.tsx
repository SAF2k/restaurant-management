"use client";

import { Tabs, TabsContent } from "@radix-ui/react-tabs";
import { useEffect, useState } from "react";
import { MenuData, getMenu } from "@/actions/get-menu";
import MenuItem from "./menu-items";

export default function MenuCollection() {
  const [menuData, setMenuData] = useState<MenuData[]>([]);

  useEffect(() => {
    const getMenuData = async () => {
      const menuData = await getMenu();
      setMenuData(menuData);
    };
    getMenuData();
  }, []);
  return (
    <div className="grid gap-4 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-5 p-2 md:p-4 lg:p-10">
      {menuData.map((menu) => (
        <Tabs defaultValue="overview" className="space-y-4" key={menu.menu_id}>
          <TabsContent value="overview" className="space-y-4">
            <div className="">
              <MenuItem key={menu.menu_id} menu={menu} />
            </div>
          </TabsContent>
        </Tabs>
      ))}
    </div>
  );
}
