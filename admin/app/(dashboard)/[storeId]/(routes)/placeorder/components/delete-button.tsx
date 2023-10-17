import { Button } from "@/components/ui/button";
import { useCartStore } from "@/hooks/use-cart";
import { Trash } from "lucide-react";

interface Props {
  ID: string;
}

export function DeleteButton({ ID }: Props) {
  const { remove } = useCartStore();

  return (
    <Button
        onClick={() => {
        remove(ID);
      }}
    >
      <Trash size={18} />
    </Button>
  );
}

export function DeleteAllButton() {
  const { removeAll } = useCartStore();

  return (
    <Button
    variant={"destructive"}
      onClick={() => {
        removeAll();
      }}
    >
      <span>Remove All</span>
      <Trash size={18} />
    </Button>
  );
}
