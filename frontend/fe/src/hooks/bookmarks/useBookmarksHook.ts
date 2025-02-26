import { BookmarkType } from "@/app/types/bookmark";
import axiosInstance from "@/lib/axiosInstance"
import { useEffect, useState } from "react";
import { useErrorHook } from "../error/useErrorHook";
import { AxiosError } from "axios";
import { toast } from "react-toastify";
import { useAppSelector } from "@/store/hook";
import { useRouter } from "next/navigation";

export const useBookmarksHook = () => {
    const { instance } = axiosInstance();
    const [bookmarks, setBookmarks] = useState<BookmarkType[] | null>(null);
    const { switchErrorHandling } = useErrorHook();
    const reload = useAppSelector(state => state.slice.reload);
    const router = useRouter();

    useEffect(() => {
        const fetchData = async () => {
            try {
                const { data, status } = await instance.get<BookmarkType[]>(
                    "/api/bookmarks",
                    { withCredentials: true }
                );
                if (status === 200) setBookmarks(data);
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
        bookmarks,
        router
    };
}