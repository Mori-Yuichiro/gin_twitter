"use client"

import Follow from "@/components/follow/Follow";
import Loading from "@/components/load/Loading";
import { useFollowersHook } from "@/hooks/follow/useFollowersHook";
import { memo } from "react";
import { ToastContainer } from "react-toastify";

const FollowersPage = memo(() => {
    const {
        profile,
        followers,
        router
    } = useFollowersHook();

    return (
        <>
            {(profile && followers) ? (
                <>
                    <div className="flex gap-x-4 items-center p-2 border-b border-gray-300">
                        <div
                            className="cursor-pointer"
                            onClick={() => router.back()}
                        >
                            <svg xmlns="http://www.w3.org/2000/svg" width="15" height="15" viewBox="0 0 512 512"><path fill="currentColor" d="M213.3 205.3v-128L0 248l213.3 170.7v-128H512v-85.4z" /></svg>
                        </div>
                        <h1 className="font-bold text-lg">{profile.displayName ? profile.displayName : profile.name}</h1>
                    </div>
                    {followers.map(follower => (
                        <div key={follower.id}>
                            <Follow follow={follower} />
                        </div>
                    ))}
                    <ToastContainer />
                </>
            ) : (
                <Loading />
            )}
        </>
    )
});
export default FollowersPage;