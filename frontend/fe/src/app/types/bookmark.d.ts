import { TweetType } from "./tweet";

export type BookmarkType = {
    id: number;
    userId: number;
    tweetId: number;
    createdAt: string;
    updatedAt: string;
    tweet: TweetType;
}