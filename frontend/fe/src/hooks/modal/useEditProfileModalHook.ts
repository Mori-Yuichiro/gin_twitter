import { ProfileType } from "@/app/types/profile";
import axiosInstance from "@/lib/axiosInstance"
import { profilePatchSchema, ProfilePatchSchemaType } from "@/lib/validations/profile";
import { useAppDispatch, useAppSelector } from "@/store/hook";
import { toggleProfileModal, toggleReload } from "@/store/slice/slice";
import { zodResolver } from "@hookform/resolvers/zod";
import { AxiosError } from "axios";
import { useCallback, useRef, useState } from "react";
import { useForm } from "react-hook-form";
import { toast } from "react-toastify";
import { useErrorHook } from "../error/useErrorHook";
import { fileRead, fileUpload, uploadImage } from "@/lib/file";

type ProfileImagesType = {
    profileImage?: string;
    avator?: string;
}

export const useEditProfileModalHook = (profile: ProfileType) => {
    const { instance } = axiosInstance();
    const openProfileModal = useAppSelector(state => state.slice.openProfileModal);
    const reload = useAppSelector(state => state.slice.reload);
    const dispatch = useAppDispatch();
    const { switchErrorHandling } = useErrorHook();
    const [profileImage, setProfileImage] = useState<string>(profile.profileImage);
    const [avator, setAvator] = useState<string>(profile.avator);

    const profileInputRef = useRef<HTMLInputElement | null>(null);
    const avatarInputRef = useRef<HTMLInputElement | null>(null);
    const fileOnClickProfile = fileUpload(profileInputRef);
    const fileOnClickAvatar = fileUpload(avatarInputRef);

    const fileInput = async (e: React.ChangeEvent<HTMLInputElement>, setImage: React.Dispatch<React.SetStateAction<string>>) => {
        const files = Array.from(e.target.files || []);

        if (files.length > 0) {
            const imageData = await fileRead(files[0]);
            setImage(imageData);
        }
    }

    const checkFormData = (data: ProfilePatchSchemaType) => {
        if (
            data.displayName === displayName ||
            data.bio === bio ||
            data.location === location ||
            data.website === website
        ) {
            return true
        }
        return false;
    }

    const updateProfileImage = async () => {
        let newProfileImagaData: string;
        let newAvatarData: string;
        let profileImages: ProfileImagesType = {};

        if (profile.profileImage !== profileImage) {
            newProfileImagaData = await (async () => {
                const imageUrl = await uploadImage(instance, profileImage);
                return imageUrl.data.data;
            })();
            profileImages.profileImage = newProfileImagaData;
        }

        if (profile.avator !== avator) {
            newAvatarData = await (async () => {
                const imageUrl = await uploadImage(instance, avator);
                return imageUrl.data.data;
            })();
            profileImages.avator = newAvatarData;
        }

        return profileImages;
    }

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
            const profileImages: ProfileImagesType = await updateProfileImage();

            if ("profileImage" in profileImages || "avator" in profileImages) {
                if (checkFormData(data)) {
                    const { status } = await instance.put<ProfileType>(
                        `/api/users/${profile.id}/edit`,
                        { ...data, ...profileImages },
                        { withCredentials: true }
                    );
                    if (status === 200) {
                        dispatch(toggleProfileModal(!openProfileModal));
                        dispatch(toggleReload(!reload));
                    }
                } else {
                    const { status } = await instance.put<ProfileType>(
                        "/users",
                        { ...profileImages },
                        { withCredentials: true }
                    );
                    if (status === 200) {
                        dispatch(toggleProfileModal(!openProfileModal));
                        dispatch(toggleReload(!reload));
                    }
                }
            } else {
                if (checkFormData(data)) {
                    const { status } = await instance.put<ProfileType>(
                        `/api/users/${profile.id}/edit`,
                        {
                            ...data,
                            profileImage: profile.profileImage,
                            avator: profile.avator
                        },
                        { withCredentials: true }
                    );
                    if (status === 200) {
                        dispatch(toggleProfileModal(!openProfileModal));
                        dispatch(toggleReload(!reload));
                    }
                }
            }
        } catch (err) {
            if (err instanceof AxiosError) {
                toast(switchErrorHandling(err.response?.data));
            } else if (err instanceof Error) {
                toast(err.message);
            }
        }
    }, [profileImage, avator]);

    return {
        setOpenProfileModal: () => dispatch(toggleProfileModal(!openProfileModal)),
        profileImage,
        avator,
        profileInputRef,
        avatarInputRef,
        fileOnClickProfile,
        fileOnClickAvatar,
        fileInputProfile: (e: React.ChangeEvent<HTMLInputElement>) => fileInput(e, setProfileImage),
        fileInputAvatar: (e: React.ChangeEvent<HTMLInputElement>) => fileInput(e, setAvator),
        register,
        handleSubmit,
        errors,
        onClickEditProfile
    }
}