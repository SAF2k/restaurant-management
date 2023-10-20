"use client";

import { useParams, useRouter } from "next/navigation";
import { Plus } from "lucide-react";
import { format } from "date-fns";
import { useEffect, useState } from "react";

import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { DataTable } from "@/components/ui/data-table";
import { Heading } from "@/components/ui/heading";
import { MenuData, getMenus } from "@/actions/get-menu";
import columns from "./columns";

export const MenuClient = () => {
  const params = useParams();
  const router = useRouter();

  const [menuData, setMenuData] = useState<MenuData[]>([]);

  const storeId = params.storeId.toString()

  useEffect(() => {
    const fetchData = async () => {
      const menuItems: MenuData[] | undefined = (await getMenus(storeId)) ?? [];
      setMenuData(menuItems);
    };
    fetchData();
  }, [ storeId ]);

  const data: MenuData[] = menuData.map((item) => ({
    _id: item._id,
    name: item.name,
    category: item.category,
    created_at: format(new Date(item.created_at), "MMMM do, yyyy"),
  }));

  return (
    <>
      <div className="flex items-center justify-between">
        <Heading
          title={`Menu (${data.length})`}
          description="Manage menu for your store."
        />
        <Button onClick={() => router.push(`/${params.storeId}/menu/new`)}>
          <Plus className="w-4 h-4 mr-2" />
          Add New
        </Button>
      </div>
      <Separator />
      <DataTable searchKey="category" columns={columns} data={data} />
    </>
  );
};
