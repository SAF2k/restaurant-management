"use client";

import { useEffect, useState } from "react";

import { MenuData, getMenuById } from "@/actions/get-menu";
import { MenuForm } from "./components/menu-form";
import toast from "react-hot-toast";

const MenuPage = ({ params }: { params: { menuId: string } }) => {
  const [menu, setMenu] = useState<MenuData | null>(null);

  useEffect(() => {
    const fetchMenu = async () => {
      try {
        if (params.menuId == "new") {
          setMenu(null);
        } else {
          const menuData: MenuData = await getMenuById(params.menuId);
          setMenu(menuData);
        }
      } catch (error: any) {
        toast.error("Something went wrong.");
      }
    };
    fetchMenu();
  }, [params.menuId]);

  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <MenuForm initialData={menu} />
      </div>
    </div>
  );
};

export default MenuPage;
