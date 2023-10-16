"use client";

import { SkeletonLoader } from "./components/skeleton";

const Loading = () => {
  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <SkeletonLoader />
      </div>
    </div>
  );
};

export default Loading;
