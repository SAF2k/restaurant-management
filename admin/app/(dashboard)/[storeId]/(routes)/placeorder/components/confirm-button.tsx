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

export default function ConfirmButton({
  tables,
  activate,
}: {
  tables: TableData[];
  activate: boolean;
}) {
  return (
    <>
      <Dialog>
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
                      <SelectItem key={item._id} value={item._id}>
                        {item.table_number}
                      </SelectItem>
                    ))}
                  </SelectGroup>
                </SelectContent>
              </Select>
            </div>
          </div>
          <DialogFooter>
            <Button type="submit">Place Order</Button>
          </DialogFooter>
        </DialogContent>
      </Dialog>
    </>
  );
}
