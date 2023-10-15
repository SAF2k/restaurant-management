import { MenuData, getMenu } from "@/actions/get-menu";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "../ui/tabs";
import FoodCollection from "@/app/(dashboard)/[storeId]/(routes)/food/components/food-collection";

export default function CategoryTab({data}: {data: MenuData[]}) {

  return (
    <>
      <Tabs defaultValue="special" className="flex flex-row">
        <TabsList className="flex flex-col gap-2 w-60 h-fit py-4 mr-5">
          {data.map((item) => (
            <TabsTrigger
              value={item.category}
              key={item._id}
              className="w-full h-12 text-md"
            >
              <>{item.category}</>
            </TabsTrigger>
          ))}
        </TabsList>
        {data.map((item) => (
          <TabsContent value={item.category} key={item._id}>
            <FoodCollection id={item._id} />
          </TabsContent>
        ))}
      </Tabs>
    </>
  );
}
