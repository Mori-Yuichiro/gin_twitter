import { CommentType } from "./comment";
import { FavoriteType } from "./favorite";
import { RetweetType } from "./retweet";
import { TweetType } from "./tweet";
import { UserType } from "./user";

export type ProfileType = UserType &
{ tweets: TweetType[] } &
{ comments: CommentType[] } &
{ retweets: RetweetType[] } &
{ favorites: FavoriteType[] } &
{ followers: RelationshipType[] } &
{ followeds: RelationshipType[] };