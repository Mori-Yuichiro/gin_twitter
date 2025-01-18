import { cookies } from "next/headers";
import { NextRequest, NextResponse } from "next/server"

export const middleware = async (req: NextRequest) => {
    const cookie = await cookies();
    const token = cookie.get("token");
    const isAuth = !!token;
    const authPage = req.nextUrl.pathname === "/";

    if (authPage) {
        if (isAuth) {
            return NextResponse.redirect(new URL("/top", req.url));
        }
        return null;
    }

    if (!isAuth) {
        return NextResponse.redirect(new URL("/", req.url));
    }
}

export const config = {
    matcher: [
        "/",
        "/top",
        "/tweets\/([0-9]+)",
        // "/profile\/([0-9]+)",
        // "/bookmarks",
        // "/rooms",
        // "/rooms\/([0-9]+)",
        // "/notifications"
    ]
}