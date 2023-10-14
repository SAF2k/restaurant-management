import Image from "next/image";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "../ui/tabs";
import { Separator } from "../ui/separator";
import { SidebarNav } from "./sidebar-nav";

const sidebarNavItems = [
  {
    title: "Profile",
    href: "/examples/forms",
  },
  {
    title: "Account",
    href: "/examples/forms/account",
  },
  {
    title: "Appearance",
    href: "/examples/forms/appearance",
  },
  {
    title: "Notifications",
    href: "/examples/forms/notifications",
  },
  {
    title: "Display",
    href: "/examples/forms/display",
  },
];

export default function CategoryTab() {
  return (
    <>
      <Tabs defaultValue="offers" className="flex flex-row gap-10">
        <TabsList className="flex flex-col gap-2 w-60 h-auto py-4">
          <TabsTrigger className="w-56 h-12 text-md" value="offers">
            Offers
          </TabsTrigger>
          <TabsTrigger className="w-56 h-12 text-md" value="pasta">
            Pasta
          </TabsTrigger>
          <TabsTrigger className="w-56 h-12 text-md" value="milk_shake">
            Milk Shake
          </TabsTrigger>
          <TabsTrigger className="w-56 h-12 text-md" value="soft_drink">
            Soft Drink
          </TabsTrigger>
          <TabsTrigger className="w-56 h-12 text-md" value="burger">
            Burger
          </TabsTrigger>
          <TabsTrigger className="w-56 h-12 text-md" value="password">
            Password
          </TabsTrigger>
        </TabsList>
        <TabsContent value="offers">This is a Offer section</TabsContent>
        <TabsContent value="pasta">This is a Pasta Section.</TabsContent>
        <TabsContent value="milk_shake">
          This is a Milk Shake section
        </TabsContent>
        <TabsContent value="soft_drink">
          This is a Soft Drink section
        </TabsContent>
        <TabsContent value="burger">This is a Burger section</TabsContent>
        <TabsContent value="password">Change your password here.</TabsContent>
      </Tabs>
    </>
  );
}
