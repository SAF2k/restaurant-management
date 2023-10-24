"use client";

import { useParams, useRouter } from "next/navigation";
import { Plus } from "lucide-react";
import { format } from "date-fns";

import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { columns } from "./columns";
import { DataTable } from "@/components/ui/data-table";
import { Heading } from "@/components/ui/heading";
import { TableData, getAllTable } from "@/actions/get-table";
import { useEffect, useState } from "react";

export const TableClient = () => {
  const params = useParams();
  const router = useRouter();
  const storeId = params.storeId.toString();

  const [tables, setTables] = useState<TableData[]>([]);

  useEffect(() => {
    const fetchTableData = async () => {
      const tableItems: TableData[] = (await getAllTable({ storeId })) ?? [];

      setTables(tableItems);
    };
    fetchTableData();
  }, [storeId]);

  const data = tables.map((item) => ({
    id: item.table_id,
    tableNumber: item.table_number,
    numberOfGuests: item.number_of_guests,
    createdAt: format(new Date(item.created_at), "MMMM do, yyyy"),
  }));

  return (
    <>
      <div className="flex items-center justify-between">
        <Heading
          title={`Table (${data.length})`}
          description="Manage table for your store."
        />
        <Button onClick={() => router.push(`/${params.storeId}/table/new`)}>
          <Plus className="w-4 h-4 mr-2" />
          Add New
        </Button>
      </div>
      <Separator />
      <DataTable searchKey="tableNumber" columns={columns} data={data} />
    </>
  );
};
