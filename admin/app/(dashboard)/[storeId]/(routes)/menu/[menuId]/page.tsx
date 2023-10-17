import { getMenuById } from "@/actions/get-menu";
import { MenuForm } from "./components/menu-form";

const MenuPage = async ({ params }: { params: { menuId: string } }) => {
  const menuData = async () => {
    if (params.menuId === "new") return null;
    const data = await getMenuById(params.menuId);
    return data;
  };
  const menu = await menuData();
  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <MenuForm initialData={menu} />
      </div>
    </div>
  );
};

export default MenuPage;
