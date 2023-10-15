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
    accessorKey: "isArchived",
    header: "Archived",
  },
  {
    accessorKey: "isFeatured",
    header: "Featured",
  },
  {
    accessorKey: "price",
    header: "Price",
  },
  {
    accessorKey: "menu",
    header: "Menu",
  },
  {
    id: "action",
    cell: ({ row }) => <CellAction data={row.original} />,
    header: "Action",
  },
];
