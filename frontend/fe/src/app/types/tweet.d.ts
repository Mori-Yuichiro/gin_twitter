import { UserType } from "./user";

export type TweetType = {
    id: number;
    content: string;
    userId: number;
    createdAt: string;
    updatedAt: string;
    user: UserType
}