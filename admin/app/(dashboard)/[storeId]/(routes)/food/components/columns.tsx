"use client";

import { ColumnDef } from "@tanstack/react-table";
import { CellAction } from "./data-table";
import { FoodData } from "@/actions/get-food";


export const columns: ColumnDef<FoodData>[] = [
  {
    accessorKey: "name",
    header: "Name",
  },
  {
    accessorKey: "price",
    header: "Price",
  },
  {
    accessorKey: "menu_name",
    header: "Menu",
  },
  {
    accessorKey: "created_at",
    header: "Created At",
  },
  {
    id: "action",
    cell: ({ row }) => <CellAction data={row.original} />,
    header: "Action",
  },
];
