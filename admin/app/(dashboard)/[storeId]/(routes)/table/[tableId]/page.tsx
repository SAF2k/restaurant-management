import { TableForm } from "./components/table-form";
import { getTableById } from "@/actions/get-table";

const FoodPage = async ({ params }: { params: { tableId: string } }) => {
  const tableData = async () => {
    try {
      if (params.tableId === "new") return null;
      const data = await getTableById(params.tableId);
      return data;
    } catch (error) {
      console.log(error);
    }
  };

 const table = await tableData()

  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <TableForm initialData={table} />
      </div>
    </div>
  );
};

export default FoodPage;
