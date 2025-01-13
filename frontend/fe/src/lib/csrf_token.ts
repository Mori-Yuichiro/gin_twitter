import { CsrfTokenType } from "@/app/types/csrf";
import axios from "axios";

export const getCsrfToken = async () => {
    const { data } = await axios.get<CsrfTokenType>(
        "http://localhost:8080/api/csrf",
        { withCredentials: true }
    );
    return data.csrf_token;
}