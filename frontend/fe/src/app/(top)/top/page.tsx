"use client"

import Button from "@/components/button/Button";
import Loading from "@/components/load/Loading";
import Tweet from "@/components/tweet/Tweet";
import { useToppageHook } from "@/hooks/top/useToppageHook";
import { ToastContainer } from "react-toastify";

export default function TopPage() {
    const {
        tweets,
        register,
        handleSubmit,
        errors,
        onClickPostTweet
    } = useToppageHook()

    return (
        <>
            {tweets ? (
                <>
                    <div>
                        <ul className="flex flex-wrap text-sm font-medium text-center text-gray-500 border-b border-gray-200 dark:border-gray-700 dark:text-gray-400">
                            <li className="w-1/2">
                                <a href="#" aria-current="page" className="inline-block p-4 text-blue-600 bg-gray-100 rounded-t-lg active dark:bg-gray-800 dark:text-blue-500">For you</a>
                            </li>
                            <li className="w-1/2">
                                <a href="#" className="inline-block p-4 rounded-t-lg hover:text-gray-600 hover:bg-gray-50 dark:hover:bg-gray-800 dark:hover:text-gray-300">Following</a>
                            </li>
                        </ul>
                        {errors.content && <p className="text-center text-red-500">{errors.content.message}</p>}
                        <div className="flex gap-x-2 border-b border-gray-300 p-3 justify-center">
                            <input
                                id="content"
                                type="text"
                                placeholder="What is happening"
                                className="w-3/4 px-2 py-1 border-gray-300 border rounded-full"
                                {...register("content")}
                            />
                            <Button
                                className="text-white bg-cyan-400 px-3 py-1 rounded-full"
                                onClick={handleSubmit(onClickPostTweet)}
                            >Post</Button>
                        </div>
                        {tweets.map(tweet => (
                            <div key={tweet.id}>
                                <Tweet tweet={tweet} />
                            </div>
                        ))}
                    </div>
                    <ToastContainer />
                </>
            ) : (
                <Loading />
            )}
        </>
    );
}