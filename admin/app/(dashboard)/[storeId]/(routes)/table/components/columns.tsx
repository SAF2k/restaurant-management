"use client";

import { ColumnDef } from "@tanstack/react-table";
import { CellAction } from "./data-table";
import { TableData } from "@/actions/get-table";

// This type is used to define the shape of our data.
// You can use a Zod schema here if you want.
export type TableColumn = {
  id: string;
  tableNumber: number;
  numberOfGuests: number;
  createdAt: string;
};

export const columns: ColumnDef<TableColumn>[] = [
  {
    accessorKey: "tableNumber",
    header: "Table Number",
  },
  {
    accessorKey: "numberOfGuests",
    header: "Number of Guests",
  },
  {
    accessorKey: "createdAt",
    header: "Date",
  },
  {
    id: "action",
    cell: ({ row }) => <CellAction data={row.original} />,
    header: "Action",
  },
];
