import { CommentType } from "@/app/types/comment";
import Link from "next/link";
import { memo } from "react";

const Comment = memo(({ comment }: { comment: CommentType }) => {
    return (
        <div className="border-b border-gray-300 px-3 py-2">
            <div className="flex gap-x-3 justify-between">
                <div className="bg-slate-400 w-8 h-8 rounded-full">
                    {comment.user.avator && <img className="w-full h-full rounded-full" src={comment.user.avator} alt="icon" />}
                </div>
                <div className="w-11/12">
                    <p>{comment.user.displayName ? comment.user.displayName : comment.user.name}</p>
                    <p>{comment.comment}</p>
                </div>
            </div>
        </div>
    );
})

export default Comment;