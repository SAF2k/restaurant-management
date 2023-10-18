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

export interface FoodClientProps {
  _id: string;
  name: string;
  price: number;
  food_id: string;
  menu_name: string;
  created_at: string;
}

export const FoodClient = () => {
  const params = useParams();
  const router = useRouter();

  const [foods, setFoods] = useState<FoodClientProps[]>([]);

  useEffect(() => {
    const fetchMenuData = async () => {
      const foodItems: FoodClientProps[] | undefined =
        (await getAllFood()) ?? [];

      setFoods(foodItems);
    };
    fetchMenuData();
  }, []);

  console.log(foods);
  

  const data: FoodClientProps[] = foods.map((item) => ({
    _id: item._id, // Use the correct property name
    name: item.name,
    price: item.price,
    food_id: item.food_id,
    menu_name: item.menu_name,
    created_at: format(new Date(item.created_at), "MMMM do, yyyy"),
  }));

  console.log(foods);
  

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
