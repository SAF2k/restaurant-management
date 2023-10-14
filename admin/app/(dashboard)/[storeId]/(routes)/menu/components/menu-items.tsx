import { MenuData } from "@/actions/get-menu";
import { Card, CardContent } from "@/components/ui/card";

interface MenuItemProps {
  menu: MenuData;
}

const MenuItem: React.FC<MenuItemProps> = ({ menu }) => (
  <div key={menu.menu_id}>
    <Card>
      <CardContent className="p-4 text-center">
        <div className="text-2xl font-bold">{menu.category}</div>
        <p className="text-xs text-muted-foreground">${menu.name}</p>
      </CardContent>
    </Card>
  </div>
);

export default MenuItem;
