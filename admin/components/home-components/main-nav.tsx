import Link from "next/link";

import { cn } from "@/lib/utils";

export function MainNav({
  className,
  ...props
}: React.HTMLAttributes<HTMLElement>) {
  return (
    <nav
      className={cn("flex items-center space-x-4 lg:space-x-6", className)}
      {...props}
    >
      <Link
        href="/id"
        className="text-sm font-medium transition-colors hover:text-primary"
      >
        Overview
      </Link>
      <Link
        href="/id/menu"
        className="text-sm font-medium text-muted-foreground transition-colors hover:text-primary"
      >
        Menu
      </Link>
      <Link
        href="/id/food"
        className="text-sm font-medium text-muted-foreground transition-colors hover:text-primary"
      >
        Food
      </Link>
      <Link
        href="/id/table"
        className="text-sm font-medium text-muted-foreground transition-colors hover:text-primary"
      >
        Table
      </Link>
      <Link
        href="/id/orders"
        className="text-sm font-medium text-muted-foreground transition-colors hover:text-primary"
      >
        Orders
      </Link>
    </nav>
  );
}
