import { z } from "zod";

export const profilePatchSchema = z.object({
    displayName: z.string()
        .min(1, { message: "displayNameは1文字以上入力してください。" }),
    bio: z.string().optional(),
    location: z.string().optional(),
    website: z.string().optional()
});

export type ProfilePatchSchemaType = z.infer<typeof profilePatchSchema>;