import { TweetType } from "./tweet";

export type FavoriteType = {
    id: number;
    userId: number;
    tweetId: number;
    createdAt: string;
    updatedAt: string;
    tweet: TweetType;
}