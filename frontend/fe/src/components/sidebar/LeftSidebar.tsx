"use client"

import { useLeftSidebarHook } from "@/hooks/sidebar/useLeftSidebarHook";
import Link from "next/link";
import Button from "../button/Button";

export default function LeftSidebar() {
    const {
        ITEM_LIST,
        currentUser,
        onClickLogOut
    } = useLeftSidebarHook();

    return (
        <aside className="fixed w-1/4 max-lg:w-1/6 h-screen overflow-y-auto py-10 border-r border-gray-300">
            <div className="flex flex-col gap-y-5 w-1/2 mx-auto max-lg:items-center">
                {ITEM_LIST.map((item, i) => (
                    <div key={`sidebar-${i}`}>
                        {item.disabled ? (
                            <div className="flex gap-x-3">
                                <div dangerouslySetInnerHTML={{ __html: item.icon }}></div>
                                <h1 className="text-xl max-lg:hidden">{item.label}</h1>
                            </div>
                        ) : (
                            <Link
                                href={item.href}
                                className="flex gap-x-3"
                            >
                                <div dangerouslySetInnerHTML={{ __html: item.icon }}></div>
                                <h1 className="text-xl max-lg:hidden">{item.label}</h1>
                            </Link>
                        )}
                    </div>
                ))}
                <Button className="bg-cyan-400 text-white rounded-full py-1 w-full" disabled>Post</Button>
                {currentUser && (
                    <div
                        className="flex gap-x-3 items-center cursor-pointer"
                        onClick={onClickLogOut}
                    >
                        <div className="bg-slate-400 w-8 h-8 rounded-full">
                            {currentUser?.avator && <img className="w-full h-full rounded-full" src={currentUser.avator} alt="icon" />}
                        </div>
                        {currentUser.displayName ? (
                            <p>{currentUser.displayName}</p>
                        ) : (
                            <p>{currentUser.name}</p>
                        )}
                    </div>
                )}
            </div>
        </aside>
    );
}