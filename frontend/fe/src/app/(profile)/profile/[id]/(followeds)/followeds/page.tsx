"use client"

import Follow from "@/components/follow/Follow";
import Loading from "@/components/load/Loading";
import { useFollowedsHook } from "@/hooks/follow/useFollowedsHook";
import { memo } from "react";
import { ToastContainer } from "react-toastify";

const FollowedsPage = memo(() => {
    const {
        profile,
        followeds,
        router
    } = useFollowedsHook();

    return (
        <>
            {(profile && followeds) ? (
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
                    {followeds.map(followed => (
                        <div key={followed.id}>
                            <Follow follow={followed} />
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
export default FollowedsPage;