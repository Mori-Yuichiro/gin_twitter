"use client"

import Button from "@/components/button/Button";
import Comment from "@/components/comment/Comment";
import Loading from "@/components/load/Loading";
import Tweet from "@/components/tweet/Tweet";
import { useTweetDetailHook } from "@/hooks/tweet-detail/useTweetDetailHook";
import { memo } from "react";
import { ToastContainer } from "react-toastify";

const TweetDetail = memo(() => {
    const {
        router,
        tweet,
        register,
        handleSubmit,
        errors,
        onClickPostComment
    } = useTweetDetailHook()

    return (
        <>
            {tweet ? (
                <>
                    <div className="flex items-center gap-x-3 px-3 py-2 border-b border-gray-200">
                        <div
                            className="cursor-pointer"
                            onClick={() => router.back()}
                        >
                            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 256 256"><path fill="currentColor" d="M224 128a8 8 0 0 1-8 8H59.31l58.35 58.34a8 8 0 0 1-11.32 11.32l-72-72a8 8 0 0 1 0-11.32l72-72a8 8 0 0 1 11.32 11.32L59.31 120H216a8 8 0 0 1 8 8Z" /></svg>
                        </div>
                        <h1 className="font-bold text-lg">Post</h1>
                    </div>
                    {tweet && <Tweet tweet={tweet} />}
                    {errors.comment && <p className="text-center text-red-500">{errors.comment.message}</p>}
                    <div className="border-b border-gray-300">
                        <div className="flex gap-x-3 p-3 justify-between">
                            <div className="bg-slate-400 w-8 h-8 rounded-full">
                                {tweet.user.avator && <img className="w-full h-full rounded-full" src={tweet.user.avator} alt="icon" />}
                            </div>
                            <input
                                id="comment"
                                type="text"
                                placeholder="Post your reply"
                                className="w-4/5 px-2"
                                {...register("comment")}
                            />
                            <Button
                                className="bg-cyan-400 rounded-full px-3 py-1"
                                onClick={handleSubmit(onClickPostComment)}
                            >Reply</Button>
                        </div>
                    </div>
                    {tweet.comments.map(comment => (
                        <div key={comment.id}>
                            <Comment comment={comment} />
                        </div>
                    ))}
                    <ToastContainer />
                </>
            ) : (
                <Loading />
            )}
        </>
    );
})

export default TweetDetail;