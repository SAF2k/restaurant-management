"use client";

import { ColumnDef } from "@tanstack/react-table";
import { CellAction } from "./data-table";
import { FoodClientProps } from "./client";

export const columns: ColumnDef<FoodClientProps>[] = [
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
