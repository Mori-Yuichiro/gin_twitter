import { z } from "zod";

export const profilePatchSchema = z.object({
    displayName: z.string().optional(),
    bio: z.string().optional(),
    location: z.string().optional(),
    website: z.string().optional()
});

export type ProfilePatchSchemaType = z.infer<typeof profilePatchSchema>;