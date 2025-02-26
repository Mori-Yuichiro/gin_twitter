import LeftSidebar from "@/components/sidebar/LeftSidebar"
import RightSidebar from "@/components/sidebar/RightSidebar";
import { memo } from "react";

const BookmarksLayout = memo(({ children }: { children: React.ReactNode }) => {
    return (
        <div className="flex w-full">
            <LeftSidebar />
            <main className="w-1/2 max-lg:w-5/6 mx-auto max-lg:ml-auto max-lg:mr-0">{children}</main>
            <RightSidebar />
        </div>
    );
});

export default BookmarksLayout;