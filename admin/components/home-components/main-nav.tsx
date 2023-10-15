"use client"

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
      href: `/${params.storeId}/orders`,
      label: "Orders",
      active: pathname === `/${params.storeId}/orders`,
    },
    {
      href: `/${params.storeId}/settings`,
      label: "Settings",
      active: pathname === `/${params.storeId}/settings`,
    },
  ];

  return (
    <nav className={cn("flex items-center space-x-4 lg:space-x-6", className)} {...props}>
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
