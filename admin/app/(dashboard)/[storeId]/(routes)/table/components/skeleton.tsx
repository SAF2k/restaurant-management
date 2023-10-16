import { Skeleton } from "@/components/ui/skeleton";

export function SkeletonLoader() {
  return (
    <>
      <div className="flex justify-between">
        <div>
          <Skeleton className="h-8 w-[150px] mb-1" />
          <Skeleton className="h-4 w-[220px]" />
        </div>
        <Skeleton className="h-10 w-[120px] mt-2" />
      </div>

      <div className="pt-8">
        <Skeleton className="h-10 w-[380px]" />
      </div>

      <div className="mt-2">
        <Skeleton className="h-[40vh] w-full" />
      </div>
      <div className="mt-2 flex justify-end mr-8 gap-2">
        <Skeleton className="h-8 w-24" />
        <Skeleton className="h-8 w-16" />
      </div>
    </>
  );
}
