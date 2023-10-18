"use client"

import { Skeleton } from "@/components/ui/skeleton";
import axios from "axios";
import { useEffect, useState } from "react";

export default function SkeletonLoader() {
  const [data, setData] = useState<any[]>([]);

  useEffect(() => {
    const fetchData = async () => {
      const res = await axios.get("/api/v2")
      const data = res.data;
      setData(data);
    };
    fetchData();
  }, []);

  console.log(data);
  
  return (
    <>
      <div className="flex justify-between p-8  pt-6">
        <div>
          <Skeleton className="h-8 w-[150px] mb-1" />
          <Skeleton className="h-4 w-[220px]" />
        </div>
        <Skeleton className="h-10 w-[120px] mt-2" />
      </div>
      <h2>{data}</h2>

      <div className="p-8">
        <Skeleton className="h-10 w-[380px]" />
      </div>

      <div className="mt-2 px-8">
        <Skeleton className="h-[40vh] w-full" />
      </div>
      <div className="mt-2 flex justify-end mr-8 gap-2">
        <Skeleton className="h-8 w-24" />
        <Skeleton className="h-8 w-16" />
      </div>
    </>
  );
}
