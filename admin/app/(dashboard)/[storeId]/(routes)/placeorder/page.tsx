import { MenuData, getMenus } from "@/actions/get-menu";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import CartItems from "./components/cart-items";
import ShowItems from "./components/show-items";

const PlaceOrderPage = async (params: { params: { storeId: string } }) => {
  const storeId = params.params.storeId.toString();
  const menuData: MenuData[] = await getMenus({ storeId });

  console.log("menuData", menuData);
  

  return (
    <div className="flex justify-between lg:px-16">
      {menuData.length <= 0 ? (
        <h1 className="flex h-screen justify-center items-center flex-1 text-xl sm:text-2xl lg:text-4xl font-bold">
          No menu items found
        </h1>
      ) : (
        <>
          <div className="p-8 pt-6">
            <Tabs defaultValue={menuData[0]._id}>
              <TabsList>
                {menuData.map((item) => (
                  <TabsTrigger value={item._id} key={item._id}>
                    <>{item.category}</>
                  </TabsTrigger>
                ))}
              </TabsList>
              {menuData.map((item) => (
                <TabsContent value={item._id} key={item._id}>
                  <ShowItems id={item._id} storeId={storeId} />
                </TabsContent>
              ))}
            </Tabs>
          </div>
          <CartItems />
        </>
      )}
    </div>
  );
};

export default PlaceOrderPage;
