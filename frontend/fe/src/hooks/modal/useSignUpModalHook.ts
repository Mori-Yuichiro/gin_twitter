import axiosInstance from "@/lib/axiosInstance";
import { signUpPatchSchema, SignUpPatchSchemaType } from "@/lib/validations/signup";
import { useAppDispatch, useAppSelector } from "@/store/hook"
import { toggleSignupModal } from "@/store/slice/slice";
import { zodResolver } from "@hookform/resolvers/zod";
import { useRouter } from "next/navigation";
import { useCallback } from "react";
import { useForm } from "react-hook-form";

export const useSignUpModalHook = () => {
    const openSignUpModal = useAppSelector(state => state.slice.openSignUpModal);
    const dispatch = useAppDispatch();

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
                if (resLogIn.status === 200) {
                    const router = useRouter();
                    router.push("/top");
                }
            }
        } catch (err) {
            console.log(err);
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