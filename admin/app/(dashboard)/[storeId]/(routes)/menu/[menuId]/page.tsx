import { getMenuById } from "@/actions/get-menu";
import { MenuForm } from "./components/menu-form";

const MenuPage = async ({
  params,
}: {
  params: { menuId: string; storeId: string };
}) => {
  const id = params.menuId.toString();
  const storeId = params.storeId.toString();
  const menu =
    params.menuId === "new" ? null : await getMenuById({ id, storeId });

  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <MenuForm initialData={menu} />
      </div>
    </div>
  );
};

export default MenuPage;
