import NavBar from "@/components/navbar";
import AuthProvider from "@/hooks/use-auth";
// import { auth } from "@clerk/nextjs";
import { redirect } from "next/navigation";

export default async function DashboardLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  // const { userId } = auth();

  // if (!userId) {
  //   redirect("/sign-in");
  // }

  return (
    // <AuthProvider>
    // </AuthProvider>
    <>
      <NavBar />
      {children}
    </>
  );
}
