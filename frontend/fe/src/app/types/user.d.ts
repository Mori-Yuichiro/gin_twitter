import { RelationshipType } from "./relationship";

export type UserType = {
    id: number;
    name: string;
    email: string;
    password: string;
    avator: string;
    displayName: string;
    profileImage: string;
    bio: string;
    location: string;
    website: string;
    createdAt: string;
    updatedAt: string;
    // followers: RelationshipType[];
    // followeds: RelationshipType[];
}