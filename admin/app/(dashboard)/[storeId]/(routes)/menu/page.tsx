import { MenuClient } from "./components/client";

const MenusPage = () => {

  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <MenuClient />
      </div>
    </div>
  );
};

export default MenusPage;
