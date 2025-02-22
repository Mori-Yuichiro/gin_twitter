import axiosInstance from "@/lib/axiosInstance"
import { tweetPatchSchema, TweetPatchSchemaType } from "@/lib/validations/tweet";
import { useAppDispatch, useAppSelector } from "@/store/hook";
import { toggleOpenTweetModal, toggleReload } from "@/store/slice/slice";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { useErrorHook } from "../error/useErrorHook";
import { useCallback } from "react";
import { AxiosError } from "axios";
import { toast } from "react-toastify";

export const useTweetModalHook = () => {
    const { instance } = axiosInstance();
    const openTweetModal = useAppSelector(state => state.slice.openTweetModal);
    const reload = useAppSelector(state => state.slice.reload);
    const dispatch = useAppDispatch();
    const { switchErrorHandling } = useErrorHook();

    const {
        register,
        reset,
        handleSubmit,
        formState: { errors }
    } = useForm<TweetPatchSchemaType>({
        resolver: zodResolver(tweetPatchSchema)
    });

    const onClickPostTweet = useCallback(async (data: TweetPatchSchemaType) => {
        try {
            const { status } = await instance.post(
                "/api/tweets",
                data,
                { withCredentials: true }
            );
            if (status === 201) {
                dispatch(toggleReload(!reload));
                dispatch(toggleOpenTweetModal(!openTweetModal));
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

    return {
        setOpenTweetModal: () => dispatch(toggleOpenTweetModal(!openTweetModal)),
        register,
        handleSubmit,
        errors,
        onClickPostTweet
    };
}