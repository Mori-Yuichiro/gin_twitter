import axiosInstance from "@/lib/axiosInstance"
import { AxiosError } from "axios";
import { useParams, useRouter } from "next/navigation";
import { useCallback, useEffect, useState } from "react";
import { toast } from "react-toastify";
import { useErrorHook } from "../error/useErrorHook";
import { ProfileType } from "@/app/types/profile";
import { useAppDispatch, useAppSelector } from "@/store/hook";
import { toggleProfileModal, toggleReload } from "@/store/slice/slice";

export const useProfileHook = () => {
    const { instance } = axiosInstance();
    const { id } = useParams();
    const router = useRouter();
    const { switchErrorHandling } = useErrorHook();
    const [profile, setProfile] = useState<ProfileType | null>(null);
    const [tab, setTab] = useState<"posts" | "comments" | "comments" | "retweets" | "articles" | "medias" | "likes">("posts");
    const openProfileModal = useAppSelector(state => state.slice.openProfileModal);
    const reload = useAppSelector(state => state.slice.reload);
    const currentUser = useAppSelector(state => state.slice.currentUser);
    const dispatch = useAppDispatch();

    const onClickFollow = useCallback(async () => {
        try {
            const { status } = await instance.post(
                `/api/users/${id}/follow`,
                null,
                { withCredentials: true }
            );
            if (status === 201) dispatch(toggleReload(!reload));
        } catch (err) {
            if (err instanceof AxiosError) {
                toast(switchErrorHandling(err.response?.data));
            } else if (err instanceof Error) {
                toast(err.message);
            }
        }
    }, [profile])

    const onClickUnFollow = useCallback(async () => {
        try {
            const { status } = await instance.delete(
                `/api/users/${id}/follow`,
                { withCredentials: true }
            );
            if (status === 200) dispatch(toggleReload(!reload));
        } catch (err) {
            if (err instanceof AxiosError) {
                toast(switchErrorHandling(err.response?.data));
            } else if (err instanceof Error) {
                toast(err.message);
            }
        }
    }, [profile])

    useEffect(() => {
        const fetchData = async () => {
            try {
                const { data, status } = await instance.get<ProfileType>(
                    `/api/users/${id}`,
                    { withCredentials: true }
                );
                if (status === 200) setProfile(data);
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
        router,
        profile,
        currentUser,
        tab,
        setTab,
        openProfileModal,
        setOpenProfileModal: () => dispatch(toggleProfileModal(!openProfileModal)),
        onClickFollow,
        onClickUnFollow
    };
}