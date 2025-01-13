import { z } from "zod";

export const tweetPatchSchema = z.object({
    content: z.string()
        .min(1, { message: "Tweetは1文字以上にしてください" })
        .max(140, { message: "Tweetは140文字以内にしてください" })
});

export type TweetPatchSchemaType = z.infer<typeof tweetPatchSchema>;