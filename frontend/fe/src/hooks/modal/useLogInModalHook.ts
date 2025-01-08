import axiosInstance from "@/lib/axiosInstance";
import { loginPatchSchema, LogInPatchSchemaType } from "@/lib/validations/login";
import { useAppDispatch, useAppSelector } from "@/store/hook"
import { toggleLogInModal } from "@/store/slice/slice";
import { zodResolver } from "@hookform/resolvers/zod";
import { useCallback } from "react";
import { useForm } from "react-hook-form";
import { useErrorHook } from "../error/useErrorHook";
import { AxiosError } from "axios";
import { toast } from "react-toastify";
import { useRouter } from "next/navigation";

export const useLogInModalHook = () => {
    const openLogInModal = useAppSelector(state => state.slice.openLogInModal);
    const dispatch = useAppDispatch();
    const router = useRouter();

    const { instance } = axiosInstance();

    const {
        register,
        handleSubmit,
        formState: { errors }
    } = useForm<LogInPatchSchemaType>({
        resolver: zodResolver(loginPatchSchema)
    });

    const onClickLogIn = useCallback(async (data: LogInPatchSchemaType) => {
        try {
            const { status } = await instance.post(
                "/api/login",
                data,
                { withCredentials: true }
            );
            if (status === 200) {
                dispatch(toggleLogInModal(!openLogInModal));
                router.push("/top");
            }
        } catch (err) {
            const { switchErrorHandling } = useErrorHook();
            if (err instanceof AxiosError) {
                toast(switchErrorHandling(err.response?.data));
            } else if (err instanceof Error) {
                toast(err.message);
            }
        }
    }, [])

    return {
        setOpenLogInModal: () => dispatch(toggleLogInModal(!openLogInModal)),
        register,
        handleSubmit,
        errors,
        onClickLogIn
    };
}