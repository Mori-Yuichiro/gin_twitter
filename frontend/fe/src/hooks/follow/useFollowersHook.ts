import axiosInstance from "@/lib/axiosInstance"
import { AxiosError } from "axios";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import { useErrorHook } from "../error/useErrorHook";
import { toast } from "react-toastify";
import { ProfileType } from "@/app/types/profile";
import { RelationshipType } from "@/app/types/relationship";

export const useFollowersHook = () => {
    const { instance } = axiosInstance();
    const [profile, setProfile] = useState<ProfileType | null>(null);
    const [followers, setFollowers] = useState<RelationshipType[] | null>(null);
    const { id } = useParams();
    const { switchErrorHandling } = useErrorHook();
    const router = useRouter();

    useEffect(() => {
        const fetchData = async () => {
            try {
                const { data, status } = await instance.get<ProfileType>(
                    `/api/users/${id}`,
                    { withCredentials: true }
                );
                if (status === 200) {
                    setProfile(data);
                    setFollowers(data.followers);
                }
            } catch (err) {
                if (err instanceof AxiosError) {
                    toast(switchErrorHandling(err.response?.data));
                } else if (err instanceof Error) {
                    toast(err.message);
                }
            }
        }
        fetchData();
    }, [])

    return {
        profile,
        followers,
        router
    }
}