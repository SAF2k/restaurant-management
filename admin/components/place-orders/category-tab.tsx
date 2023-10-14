"use client";
import { MenuData, getMenu } from "@/actions/get-menu";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "../ui/tabs";
import { useEffect, useState } from "react";

export default function CategoryTab() {
  const category = ["Breakfast", "Lunch", "Dinner", "Snacks", "Dessert"];

  const [menuData, setMenuData] = useState<MenuData[]>([]);

  useEffect(() => {
    const getMenuData = async () => {
      const menuData = await getMenu();
      setMenuData(menuData);
    };
    getMenuData();
  }, []);

  console.log(menuData);

  return (
    <>
      <Tabs defaultValue={category[0]} className="flex flex-row gap-10">
        <TabsList className="flex flex-col gap-2 w-60 h-auto py-4">
          {category.map((item) => (
            <TabsTrigger key={item} className="w-56 h-12 text-md" value={item}>
              {item}
            </TabsTrigger>
          ))}
        </TabsList>
        {category.map((item) => (
          <TabsContent key={item} value={item}>
            This is a {item} section
          </TabsContent>
        ))}
      </Tabs>
    </>
  );
}
