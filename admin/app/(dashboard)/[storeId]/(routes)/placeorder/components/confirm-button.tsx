import { TableData } from "@/actions/get-table";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { useState } from "react";

export default function ConfirmButton({
  tables,
  activate,
}: {
  tables: TableData[];
  activate: boolean;
}) {
  const [tableId, setTableId] = useState<string>("");
  console.log(tableId);
  const onSubmit = () => {
    console.log("submit");
  };

  return (
    <>
      <Dialog>
        <form ></form>
        <DialogTrigger asChild>
          <Button disabled={activate}>Continue Order</Button>
        </DialogTrigger>
        <DialogContent className="sm:max-w-[425px]">
          <DialogHeader>
            <DialogTitle>Choose Table</DialogTitle>
            <DialogDescription>
              {` Select table according to the order.`}
            </DialogDescription>
          </DialogHeader>
          <div className="grid gap-4 py-4">
            <div className="grid grid-cols-3 items-center justify-between gap-4">
              <Label htmlFor="table-number" className="text-right">
                Table Number
              </Label>
              <Select>
                <SelectTrigger className="w-[220px]">
                  <SelectValue placeholder="Table Number" />
                </SelectTrigger>
                <SelectContent>
                  <SelectGroup>
                    {tables.map((item) => (
                      <SelectItem key={item.table_id} value={item.table_id}>
                        {item.table_number}
                      </SelectItem>
                    ))}
                  </SelectGroup>
                </SelectContent>
              </Select>
            </div>
          </div>
          <DialogFooter>
            <Button onClick={onSubmit}>Place Order</Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </>
  );
}
