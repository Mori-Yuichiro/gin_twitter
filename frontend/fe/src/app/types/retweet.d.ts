import { TweetType } from "./tweet";

export type RetweetType = {
    id: number;
    userId: number;
    tweetId: number;
    createdAt: string;
    updatedAt: string;
    tweet: TweetType;
}