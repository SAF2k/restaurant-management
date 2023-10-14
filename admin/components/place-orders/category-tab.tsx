"use client";
import { MenuData, getMenu } from "@/actions/get-menu";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "../ui/tabs";
import { useEffect, useState } from "react";
import FoodCollection from "@/app/(dashboard)/[storeId]/(routes)/food/components/food-collection";

export default function CategoryTab() {
  const [menuData, setMenuData] = useState<MenuData[]>([]);

  useEffect(() => {
    const fetchMenuData = async () => {
      const menuItems: MenuData[] = await getMenu();
      setMenuData(menuItems);
    };
    fetchMenuData();
  }, []);

  const testData = menuData.map((item) => (
    <TabsTrigger
      value={item.category}
      key={item._id}
      className="w-full h-12 text-md"
    >
      <>
        {console.log(item.category)}
        {item.category}
      </>
    </TabsTrigger>
  ));

  console.log(testData);

  return (
    <>
      <Tabs defaultValue={menuData[0]?.category} className="flex flex-row">
        <TabsList className="flex flex-col gap-2 w-60 h-fit py-4 mr-5">
          {testData}
        </TabsList>
        {menuData.map((item) => (
          <TabsContent value={item.category} key={item._id}>
            <FoodCollection id={item._id} />
          </TabsContent>
        ))}
      </Tabs>
    </>
  );
}
