import { BookmarkType } from "./bookmark";
import { CommentType } from "./comment";
import { FavoriteType } from "./favorite";
import { RetweetType } from "./retweet";
import { UserType } from "./user";

export type TweetType = {
    id: number;
    content: string;
    userId: number;
    createdAt: string;
    updatedAt: string;
    user: UserType;
    comments: CommentType[];
    retweets: RetweetType[];
    favorites: FavoriteType[];
    bookmarks: BookmarkType[];
}