import axiosInstance from "@/lib/axiosInstance";
import { signUpPatchSchema, SignUpPatchSchemaType } from "@/lib/validations/signup";
import { useAppDispatch, useAppSelector } from "@/store/hook"
import { toggleSignupModal } from "@/store/slice/slice";
import { zodResolver } from "@hookform/resolvers/zod";
import { useRouter } from "next/navigation";
import { useCallback } from "react";
import { useForm } from "react-hook-form";
import { useErrorHook } from "../error/useErrorHook";
import { AxiosError } from "axios";
import { toast } from "react-toastify";

export const useSignUpModalHook = () => {
    const openSignUpModal = useAppSelector(state => state.slice.openSignUpModal);
    const dispatch = useAppDispatch();
    const router = useRouter();

    const { instance } = axiosInstance();

    const {
        register,
        handleSubmit,
        formState: { errors }
    } = useForm<SignUpPatchSchemaType>({
        resolver: zodResolver(signUpPatchSchema)
    });

    const onClickSignUp = useCallback(async (data: SignUpPatchSchemaType) => {
        try {
            const { status } = await instance.post(
                "/api/signup",
                data,
                { withCredentials: true }
            );

            if (status === 201) {
                const { email, password } = data;

                const resLogIn = await instance.post(
                    "/api/login",
                    { email, password },
                    { withCredentials: true }
                );
                if (resLogIn.status === 200) router.push("/top");
            }
        } catch (err) {
            const { switchErrorHandling } = useErrorHook();
            if (err instanceof AxiosError) {
                toast(switchErrorHandling(err.response?.data));
            } else if (err instanceof Error) {
                toast(switchErrorHandling(err.message));
            }
        }
    }, [])

    return {
        setOpenSignUpModal: () => dispatch(toggleSignupModal(!openSignUpModal)),
        register,
        handleSubmit,
        errors,
        onClickSignUp
    };
}