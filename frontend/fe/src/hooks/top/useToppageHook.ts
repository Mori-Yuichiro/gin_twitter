import axiosInstance from "@/lib/axiosInstance"
import { useAppSelector } from "@/store/hook";
import { useEffect, useState } from "react";
import { useErrorHook } from "../error/useErrorHook";
import { AxiosError } from "axios";
import { toast } from "react-toastify";
import { TweetType } from "@/app/types/tweet";

export const useToppageHook = () => {
    const { instance } = axiosInstance();
    const [tweets, setTweets] = useState<TweetType[] | null>(null);
    const reload = useAppSelector(state => state.slice.reload);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const { data, status } = await instance.get<TweetType[]>(
                    "/api/tweets",
                    { withCredentials: true }
                );

                if (status === 200) setTweets(data);
            } catch (err) {
                const { switchErrorHandling } = useErrorHook();
                if (err instanceof AxiosError) {
                    toast(switchErrorHandling(err.response?.data));
                } else if (err instanceof Error) {
                    toast(err.message);
                }
            }
        }
        fetchData();
    }, [reload])

    return {
        tweets
    };
}