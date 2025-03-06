"use client"

import { RelationshipType } from "@/app/types/relationship";
import { usePathname } from "next/navigation";
import { memo } from "react";

const Follow = memo(({ follow }: { follow: RelationshipType }) => {
    const url = usePathname();
    console.log(url.match("/followeds"));

    return (
        <div className="border-b border-gray-300 px-3 py-2">
            <div className="flex gap-x-3 justify-between">
                {(url.match("/followeds")) ? (
                    <>
                        <div className="bg-slate-400 w-8 h-8 rounded-full">
                            {follow.followed.avator && <img className="w-full h-full rounded-full" src={follow.followed.avator} alt="icon" />}
                        </div>
                        <div className="w-11/12">
                            <p>{follow.followed.displayName ? follow.followed.displayName : follow.followed.name}</p>
                            <p>{follow.followed.bio}</p>
                        </div>
                    </>
                ) : (url.match("/followers")) ? (
                    <>
                        <div className="bg-slate-400 w-8 h-8 rounded-full">
                            {follow.follower.avator && <img className="w-full h-full rounded-full" src={follow.follower.avator} alt="icon" />}
                        </div>
                        <div className="w-11/12">
                            <p>{follow.follower.displayName ? follow.follower.displayName : follow.follower.name}</p>
                            <p>{follow.follower.bio}</p>
                        </div>
                    </>
                ) : <></>}
            </div>
        </div>
    );
});
export default Follow;