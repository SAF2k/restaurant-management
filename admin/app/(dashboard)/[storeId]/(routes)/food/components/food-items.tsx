import { FoodData } from "@/actions/get-food";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

interface FoodItemProps {
  food: FoodData;
}

const FoodItem: React.FC<FoodItemProps> = ({ food }) => (
  <div key={food.food_id}>
    <Card>
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        {/* <img src={food.food_image} height="fit" width="fit" alt={food.name} /> */}
      </CardHeader>
      <CardContent>
        <div className="text-2xl font-bold">{food.name}</div>
        <p className="text-xs text-muted-foreground">${food.price}</p>
      </CardContent>
    </Card>
  </div>
);

export default FoodItem;
