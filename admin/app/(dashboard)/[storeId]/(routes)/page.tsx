import OverviewContent from "@/components/home-components/overview-content";
import PlaceOrder from "@/components/place-orders/place-content";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";

export default function HomePage() {
  return (
    <>
      <div className="flex-row md:flex-col md:flex">
        <div className="flex-1 space-y-4 p-2 md:p-8 pt-6">
          <div className="flex flex-col md:flex-row items-center justify-between space-y-2">
            <h2 className="text-3xl font-bold tracking-tight">Dashboard</h2>
            {/* <div className="flex items-center space-x-2">
              <CalendarDateRangePicker />
              <Button>Download</Button>
            </div> */}
          </div>
          <Tabs defaultValue="overview" className="space-y-4">
            <TabsList>
              <TabsTrigger value="overview">Overview</TabsTrigger>
              <TabsTrigger value="place_order">Place Order</TabsTrigger>
            </TabsList>
            <TabsContent value="overview" className="space-y-4">
              <OverviewContent />
            </TabsContent>
            <TabsContent value="place_order" className="space-y-4">
              <PlaceOrder />
            </TabsContent>
          </Tabs>
        </div>
      </div>
    </>
  );
}
