"use client"

import { useEffect, useState } from "react";
import CategoryTab from "./category-tab";
import { MenuData, getMenu } from "@/actions/get-menu";

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
        <div className="flex">
            <CategoryTab data={menuData}/>
        </div>
        </>
    )
}