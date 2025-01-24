import axiosInstance from "@/lib/axiosInstance"
import { AxiosError } from "axios";
import { useErrorHook } from "../error/useErrorHook";
import { useAppDispatch, useAppSelector } from "@/store/hook";
import { toggleReload } from "@/store/slice/slice";
import { usePathname } from "next/navigation";
import { useCallback } from "react";

export const useTweetHook = (id: number) => {
    const { instance } = axiosInstance();
    const { switchErrorHandling } = useErrorHook();
    const reload = useAppSelector(state => state.slice.reload);
    const currentUser = useAppSelector(state => state.slice.currentUser);
    const dispatch = useAppDispatch();
    const pathName = usePathname();

    const onClickDeleteTweet = useCallback(async () => {
        try {
            if (window.confirm("こちらのツイートを削除します。よろしいですか？")) {
                const { status } = await instance.delete(
                    `/api/tweets/${id}`,
                    { withCredentials: true }
                );
                if (status === 204) dispatch(toggleReload(!reload));
            }
        } catch (err) {
            if (err instanceof AxiosError) {
                console.error(switchErrorHandling(err.response?.data));
            } else if (err instanceof Error) {
                console.error(err.message);
            }
        }
    }, [reload]);

    const onClickCreateRetweet = useCallback(async () => {
        try {
            const { status } = await instance.post(
                `/api/tweets/${id}/retweet`,
                undefined,
                { withCredentials: true }
            );
            if (status === 201) dispatch(toggleReload(!reload));
        } catch (err) {
            if (err instanceof AxiosError) {
                console.error(switchErrorHandling(err.response?.data));
            } else if (err instanceof Error) {
                console.error(err.message);
            }
        }
    }, [reload]);

    const onClickDeleteRetweet = useCallback(async () => {
        try {
            const { status } = await instance.delete(
                `/api/tweets/${id}/retweet`,
                { withCredentials: true }
            );
            if (status === 200) dispatch(toggleReload(!reload));
        } catch (err) {
            if (err instanceof AxiosError) {
                console.error(switchErrorHandling(err.response?.data));
            } else if (err instanceof Error) {
                console.error(err.message);
            }
        }
    }, [reload])

    return {
        onClickDeleteTweet,
        currentUser,
        pathName,
        onClickCreateRetweet,
        onClickDeleteRetweet
    };
}