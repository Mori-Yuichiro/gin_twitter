import { TweetType } from "@/app/types/tweet";
import axiosInstance from "@/lib/axiosInstance";
import { AxiosError } from "axios";
import { useParams, useRouter } from "next/navigation"
import { useEffect, useState } from "react";
import { toast } from "react-toastify";
import { useErrorHook } from "../error/useErrorHook";

export const useTweetDetailHook = () => {
    const router = useRouter();
    const [tweet, setTweet] = useState<TweetType | null>(null);
    const { instance } = axiosInstance();
    const { switchErrorHandling } = useErrorHook();
    const { id } = useParams();

    useEffect(() => {
        const fetchData = async () => {
            try {
                const { data, status } = await instance.get<TweetType>(
                    `/api/tweets/${id}`,
                    { withCredentials: true }
                );
                if (status === 200) setTweet(data);
            } catch (err) {
                if (err instanceof AxiosError) {
                    toast(switchErrorHandling(err.response?.data));
                } else if (err instanceof Error) {
                    toast(err.message);
                }
            }
        };
        fetchData();
    }, [])

    return {
        router,
        tweet
    };
}