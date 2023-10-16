"use client"

import { TableClient } from "./components/client";

const TablesPage = () => {
  return (
    <div className="flex-col">
      <div className="flex-1 space-y-4 p-8 pt-6">
        <TableClient />
      </div>
    </div>
  );
};

export default TablesPage;
