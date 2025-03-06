import { UserType } from "./user";

export type RelationshipType = {
    id: number;
    followerId: number;
    followedId: number;
    createdAt: string;
    updatedAt: string;
    follower: UserType;
    followed: UserType;
}