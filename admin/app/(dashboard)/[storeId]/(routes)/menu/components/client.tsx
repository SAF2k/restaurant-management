"use client";

import { useParams, useRouter } from "next/navigation";
import { Plus } from "lucide-react";
import { format } from "date-fns";

import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { columns } from "./columns";
import { DataTable } from "@/components/ui/data-table";
import { Heading } from "@/components/ui/heading";
import { useEffect, useState } from "react";
import { MenuData, getMenu } from "@/actions/get-menu";

export const MenuClient = () => {
  const params = useParams();
  const router = useRouter();

  const [menuData, setMenuData] = useState<MenuData[]>([]);

  useEffect(() => {
    const fetchMenuData = async () => {
      const menuItems: MenuData[] = await getMenu();
      setMenuData(menuItems);
    };
    fetchMenuData();
  }, []);

  const data = menuData.map((item) => ({
    id: item._id,
    label: item.category,
    createdAt: format(new Date(item.created_at), "MMMM do, yyyy"),
  }));

  return (
    <>
      <div className="flex items-center justify-between">
        <Heading
          title={`Menu (${data.length})`}
          description="Manage menu for your store."
        />
        <Button
          onClick={() => router.push(`/${params.storeId}/menu/new`)}
        >
          <Plus className="w-4 h-4 mr-2" />
          Add New
        </Button>
      </div>
      <Separator />
      <DataTable searchKey="label" columns={columns} data={data} />
    </>
  );
};
