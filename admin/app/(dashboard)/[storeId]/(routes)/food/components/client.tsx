"use client";

import { useParams, useRouter } from "next/navigation";
import { Plus } from "lucide-react";
import { useEffect, useState } from "react";
import { format } from "date-fns";

import { Button } from "@/components/ui/button";
import { Heading } from "@/components/ui/heading";
import { Separator } from "@/components/ui/separator";
import { DataTable } from "@/components/ui/data-table";
import { columns } from "./columns";
import { FoodData, getAllFood } from "@/actions/get-food";

export const FoodClient = () => {
  const params = useParams();
  const router = useRouter();
  const storeId = params.storeId.toString();

  const [foods, setFoods] = useState<FoodData[]>([]);

  useEffect(() => {
    const fetchMenuData = async () => {
      const foodItems: FoodData[] | undefined =
        (await getAllFood({ storeId })) ?? [];

      setFoods(foodItems);
    };
    fetchMenuData();
  }, [storeId]);

  const data: FoodData[] = foods.map((item) => ({
    ID: item.ID,
    name: item.name,
    price: item.price,
    food_id: item.food_id,
    menu_name: item.menu_name,
    created_at: format(new Date(item.created_at), "MMMM do, yyyy"),
  }));

  return (
    <>
      <div className="flex items-center justify-between">
        <Heading
          title={`Food (${data.length})`}
          description="Manage Food for your store."
        />
        <Button onClick={() => router.push(`/${params.storeId}/food/new`)}>
          <Plus className="w-4 h-4 mr-2" />
          Add New
        </Button>
      </div>
      <Separator />
      <DataTable searchKey="name" columns={columns} data={data} />
    </>
  );
};
