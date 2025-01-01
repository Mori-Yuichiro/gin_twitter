import { CsrfTokenType } from "@/app/types/csrf";
import axiosInstance from "./axiosInstance";

export const getCsrfToken = async () => {
    const { instance } = axiosInstance();
    const { data } = await instance.get<CsrfTokenType>(
        "/api/csrf",
        { withCredentials: true }
    );
    return data.csrf_token;
}