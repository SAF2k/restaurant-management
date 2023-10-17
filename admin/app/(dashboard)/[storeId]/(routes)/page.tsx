import OverviewContent from "@/components/home-components/overview-content";
import { Button } from "@/components/ui/button";
import Link from "next/link";

export default function HomePage({ params }: { params: { storeId: string } }) {
  return (
    <>
      <div className="flex-row lg:px-8 xl:px-14 md:flex-col md:flex">
        <div className="flex-1 space-y-4 p-2 md:p-8 pt-6">
          <div className="flex flex-col md:flex-row items-center justify-between space-y-2">
            <h2 className="text-3xl font-bold tracking-tight">Dashboard</h2>
            <div className="flex items-center space-x-2">
              {/* <CalendarDateRangePicker /> */}
              <Link href={`${params.storeId}/cart`}>
                <Button>Place Orders</Button>
              </Link>
            </div>
          </div>
          <OverviewContent />
        </div>
      </div>
    </>
  );
}
