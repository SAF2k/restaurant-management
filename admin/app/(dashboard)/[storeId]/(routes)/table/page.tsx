import { TableColumn } from "./components/columns";
import { format } from "date-fns";
import { TableClient } from "./components/client";

const Tables: TableColumn[] = [
  {
    id: "1",
    label: "Table 1",
    createdAt: "2021-10-10T16:00:00.000Z",
  },
  {
    id: "2",
    label: "Table 2",
    createdAt: "2021-10-10T16:00:00.000Z",
  },
  {
    id: "3",
    label: "Table 3",
    createdAt: "2021-10-10T16:00:00.000Z",
  },
  {
    id: "4",
    label: "Table 4",
    createdAt: "2021-10-10T16:00:00.000Z",
  },
  {
    id: "5",
    label: "Table 5",
    createdAt: "2021-10-10T16:00:00.000Z",
  },
  {
    id: "6",
    label: "Table 6",
    createdAt: "2021-10-10T16:00:00.000Z",
  },
  {
    id: "7",
    label: "Table 7",
    createdAt: "2021-10-10T16:00:00.000Z",
  },
  {
    id: "8",
    label: "Table 8",
    createdAt: "2021-10-10T16:00:00.000Z",
  },
  {
    id: "9",
    label: "Table 9",
    createdAt: "2021-10-10T16:00:00.000Z",
  },
  {
    id: "10",
    label: "Table 10",
    createdAt: "2021-10-10T16:00:00.000Z",
  },
  {
    id: "1",
    label: "Table 1",
    createdAt: "2021-10-10T16:00:00.000Z",
  },
  {
    id: "2",
    label: "Table 2",
    createdAt: "2021-10-10T16:00:00.000Z",
  },
  {
    id: "3",
    label: "Table 3",
    createdAt: "2021-10-10T16:00:00.000Z",
  },
  {
    id: "4",
    label: "Table 4",
    createdAt: "2021-10-10T16:00:00.000Z",
  },
  {
    id: "5",
    label: "Table 5",
    createdAt: "2021-10-10T16:00:00.000Z",
  },
  {
    id: "6",
    label: "Table 6",
    createdAt: "2021-10-10T16:00:00.000Z",
  },
  {
    id: "7",
    label: "Table 7",
    createdAt: "2021-10-10T16:00:00.000Z",
  },
  {
    id: "8",
    label: "Table 8",
    createdAt: "2021-10-10T16:00:00.000Z",
  },
  {
    id: "9",
    label: "Table 9",
    createdAt: "2021-10-10T16:00:00.000Z",
  },
  {
    id: "10",
    label: "Table 10",
    createdAt: "2021-10-10T16:00:00.000Z",
  },
];

const TablesPage = () => {
  const formattedTables: TableColumn[] = Tables.map((item) => ({
    id: item.id,
    label: item.label,
    createdAt: format(new Date(item.createdAt), "MMMM do, yyyy"),
  }));

  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <TableClient data={formattedTables} />
      </div>
    </div>
  );
};

export default TablesPage;
