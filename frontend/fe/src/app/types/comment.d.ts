import { UserType } from "./user";

export type CommentType = {
    id: number;
    comment: string;
    userId: number;
    tweetId: number;
    createdAt: string;
    updatedAt: string;
    user: UserType
}