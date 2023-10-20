import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

export function middleware(request: NextRequest) {
  const path = request.nextUrl.pathname;

  const isProtected = path === "/login" || path === "/signup";

  const token = request.cookies.get("token")?.value || "";

  console.log("token:", token);

  if (isProtected && token) {
    return NextResponse.redirect(new URL("/profile", request.nextUrl));
  }

  if (!isProtected && !token) {
    return NextResponse.redirect(new URL("/login", request.nextUrl));
  }
}

export const config = {
  matcher: ["/", "/[id]/", "/login", "/signup"],
};
