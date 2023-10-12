import { MainNav } from "./home-components/main-nav";
import { UserNav } from "./home-components/user-nav";
import { ModeToggle } from "./theme-toggle";

export default async function NavBar () {
 return (
   <div className="border-b">
     <div className="flex h-16 items-center px-4">
       {/* <TeamSwitcher /> */}
       <MainNav className="mx-6" />
       <div className="md:ml-auto flex items-center space-x-4">
         {/* <Search /> */}
         <ModeToggle />
         <UserNav />
       </div>
     </div>
   </div>
 );
}