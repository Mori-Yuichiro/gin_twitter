import { getCsrfToken } from "@/lib/csrf_token";
import { useAppDispatch, useAppSelector } from "@/store/hook";
import { toggleSignupModal } from "@/store/slice/slice";
import axios from "axios";
import { useEffect } from "react";

export const useHomeHook = () => {
    const openSignUpModal = useAppSelector(state => state.slice.openSignUpModal);
    const dispatch = useAppDispatch();

    useEffect(() => {
        try {
            const fetchData = async () => {
                const csrf_token = await getCsrfToken();
                axios.defaults.withCredentials = true;
                axios.defaults.headers.common["X-CSRF-TOKEN"] = csrf_token;
            }
            fetchData();
        } catch (err) {
            console.log(err);
        }
    }, [])

    return {
        openSignUpModal,
        setOpenSignUpModal: () => dispatch(toggleSignupModal(!openSignUpModal))
    };
}