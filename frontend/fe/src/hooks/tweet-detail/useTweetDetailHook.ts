import { TweetType } from "@/app/types/tweet";
import axiosInstance from "@/lib/axiosInstance";
import { AxiosError } from "axios";
import { useParams, useRouter } from "next/navigation"
import { useCallback, useEffect, useState } from "react";
import { toast } from "react-toastify";
import { useErrorHook } from "../error/useErrorHook";
import { useForm } from "react-hook-form";
import { commentPatchSchema, CommentPatchSchemaType } from "@/lib/validations/comment";
import { zodResolver } from "@hookform/resolvers/zod";
import { useAppDispatch, useAppSelector } from "@/store/hook";
import { toggleReload } from "@/store/slice/slice";

export const useTweetDetailHook = () => {
    const router = useRouter();
    const [tweet, setTweet] = useState<TweetType | null>(null);
    const { instance } = axiosInstance();
    const { switchErrorHandling } = useErrorHook();
    const { id } = useParams();
    const reload = useAppSelector(state => state.slice.reload);
    const dispatch = useAppDispatch();

    const {
        register,
        handleSubmit,
        reset,
        formState: { errors }
    } = useForm<CommentPatchSchemaType>({
        resolver: zodResolver(commentPatchSchema)
    });

    const onClickPostComment = useCallback(async (data: CommentPatchSchemaType) => {
        try {
            const { status } = await instance.post(
                "/api/comment",
                {
                    ...data,
                    tweetId: Number(id)
                },
                { withCredentials: true }
            );
            if (status === 201) {
                dispatch(toggleReload(!reload));
                reset({ comment: "" });
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
    }, [reload])

    return {
        router,
        tweet,
        register,
        handleSubmit,
        errors,
        onClickPostComment
    };
}