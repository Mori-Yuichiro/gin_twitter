import { ProfileType } from "@/app/types/profile";
import axiosInstance from "@/lib/axiosInstance"
import { profilePatchSchema, ProfilePatchSchemaType } from "@/lib/validations/profile";
import { useAppDispatch, useAppSelector } from "@/store/hook";
import { toggleProfileModal, toggleReload } from "@/store/slice/slice";
import { zodResolver } from "@hookform/resolvers/zod";
import { AxiosError } from "axios";
import { useCallback } from "react";
import { useForm } from "react-hook-form";
import { toast } from "react-toastify";
import { useErrorHook } from "../error/useErrorHook";

export const useEditProfileModalHook = (profile: ProfileType) => {
    const { instance } = axiosInstance();
    const openProfileModal = useAppSelector(state => state.slice.openProfileModal);
    const reload = useAppSelector(state => state.slice.reload);
    const dispatch = useAppDispatch();
    const { switchErrorHandling } = useErrorHook();

    const {
        displayName,
        bio,
        location,
        website
    } = profile;
    const {
        register,
        handleSubmit,
        formState: { errors }
    } = useForm<ProfilePatchSchemaType>({
        resolver: zodResolver(profilePatchSchema),
        defaultValues: {
            displayName,
            bio,
            location,
            website
        }
    });

    const onClickEditProfile = useCallback(async (data: ProfilePatchSchemaType) => {
        try {
            const { status } = await instance.put<ProfileType>(
                `/api/users/${profile.id}/edit`,
                data,
                { withCredentials: true }
            );

            if (status === 200) {
                dispatch(toggleProfileModal(!openProfileModal));
                dispatch(toggleReload(!reload));
            }
        } catch (err) {
            if (err instanceof AxiosError) {
                toast(switchErrorHandling(err.response?.data));
            } else if (err instanceof Error) {
                toast(err.message);
            }
        }
    }, []);

    return {
        setOpenProfileModal: () => dispatch(toggleProfileModal(!openProfileModal)),
        register,
        handleSubmit,
        errors,
        onClickEditProfile
    }
}