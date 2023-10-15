"use client";

import { ColumnDef } from "@tanstack/react-table";
import { CellAction } from "./data-table";

// This type is used to define the shape of our data.
// You can use a Zod schema here if you want.
export type TableColumn = {
  id: string;
  label: string;
  createdAt: string;
};

export const columns: ColumnDef<TableColumn>[] = [
  {
    accessorKey: "label",
    header: "Label",
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
