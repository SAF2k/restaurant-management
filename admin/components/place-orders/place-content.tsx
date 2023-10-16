import { MenuData, getMenu } from "@/actions/get-menu";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "../ui/tabs";

export default async function PlaceOrder() {
  const menuData: MenuData[] = await getMenu();

  return (
    <>
      <div className="flex mt-8">
        <Tabs defaultValue={menuData[0].value} className="flex flex-row">
          <TabsList className="flex flex-col gap-2 sm:w-40 lg:w-60 h-fit py-4 mr-5">
            {menuData.map((item) => (
              <TabsTrigger
                value={item.category}
                key={item._id}
                className="w-full h-12 text-md"
              >
                <>{item.category}</>
              </TabsTrigger>
            ))}
          </TabsList>
          {menuData.map((item) => (
            <TabsContent value={item.category} key={item._id}>
              {item.category}
            </TabsContent>
          ))}
        </Tabs>
      </div>
    </>
  );
}
