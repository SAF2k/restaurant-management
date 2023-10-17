"use client";

import Link from "next/link";

import { cn } from "@/lib/utils";
import { useParams, usePathname } from "next/navigation";

export function MainNav({
  className,
  ...props
}: React.HTMLAttributes<HTMLElement>) {
  const pathname = usePathname();
  const params = useParams();

  const routes = [
    {
      href: `/${params.storeId}`,
      label: "Overview",
      active: pathname === `/${params.storeId}`,
    },
    {
      href: `/${params.storeId}/placeorder`,
      label: "Place Order",
      active: pathname === `/${params.storeId}/placeorder`,
    },
    {
      href: `/${params.storeId}/menu`,
      label: "Menu",
      active: pathname === `/${params.storeId}/menu`,
    },
    {
      href: `/${params.storeId}/food`,
      label: "Food",
      active: pathname === `/${params.storeId}/food`,
    },
    {
      href: `/${params.storeId}/table`,
      label: "Table",
      active: pathname === `/${params.storeId}/table`,
    },
    {
      href: `/${params.storeId}/skeleton`,
      label: "Skeleton",
      active: pathname === `/${params.storeId}/skeleton`,
    },
  ];

  return (
    <nav
      className={cn(
        "hidden sm:flex items-center space-x-4 lg:space-x-6",
        className
      )}
      {...props}
    >
      {routes.map((route) => (
        <Link
          key={route.href}
          href={route.href}
          className={cn(
            "text-sm font-medium transition-color hover:text-primary",
            route.active
              ? "text-[black] dark:text-white"
              : "text-muted-foreground"
          )}
        >
          {route.label}
        </Link>
      ))}
    </nav>
  );
}
