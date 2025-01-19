import axiosInstance from "@/lib/axiosInstance"
import { useAppDispatch, useAppSelector } from "@/store/hook";
import { useCallback, useEffect, useState } from "react";
import { useErrorHook } from "../error/useErrorHook";
import { AxiosError } from "axios";
import { toast } from "react-toastify";
import { TweetType } from "@/app/types/tweet";
import { useForm } from "react-hook-form";
import { tweetPatchSchema, TweetPatchSchemaType } from "@/lib/validations/tweet";
import { zodResolver } from "@hookform/resolvers/zod";
import { toggleReload } from "@/store/slice/slice";

export const useToppageHook = () => {
    const { instance } = axiosInstance();
    const [tweets, setTweets] = useState<TweetType[] | null>(null);
    const reload = useAppSelector(state => state.slice.reload);
    const dispatch = useAppDispatch();
    const { switchErrorHandling } = useErrorHook();

    const {
        register,
        handleSubmit,
        reset,
        formState: { errors }
    } = useForm<TweetPatchSchemaType>({
        resolver: zodResolver(tweetPatchSchema)
    })

    const onClickPostTweet = useCallback(async (data: TweetPatchSchemaType) => {
        try {
            const { status } = await instance.post(
                "/api/tweets",
                data,
                { withCredentials: true }
            );
            if (status === 201) {
                dispatch(toggleReload(!reload));
                reset({ content: "" });
            }
        } catch (err) {
            if (err instanceof AxiosError) {
                toast(switchErrorHandling(err.response?.data));
            } else if (err instanceof Error) {
                toast(err.message);
            }
        }
    }, [reload])

    useEffect(() => {
        const fetchData = async () => {
            try {
                const { data, status } = await instance.get<TweetType[]>(
                    "/api/tweets",
                    { withCredentials: true }
                );
                if (status === 200) setTweets(data);
            } catch (err) {
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
        tweets,
        register,
        handleSubmit,
        errors,
        onClickPostTweet
    };
}