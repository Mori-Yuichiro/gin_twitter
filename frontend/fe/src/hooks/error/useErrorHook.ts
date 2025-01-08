import { CsrfTokenType } from "@/app/types/csrf";
import axiosInstance from "@/lib/axiosInstance";

export const useErrorHook = () => {
    const { instance } = axiosInstance();
    const setCsrfToken = async () => {
        const { data } = await instance.get<CsrfTokenType>("/api/csrf");
        instance.defaults.headers.common['X-CSRF-TOKEN'] = data.csrf_token;
    }

    const switchErrorHandling = (msg: string) => {
        switch (msg) {
            case "missing csrf token in request header":
                // getCsrfToken();
                return "CSRFトークンがヘッダーに設定されていません";
            case "invalid csrf token":
                // getCsrfToken();
                return "CSRFトークンが間違っています";
            case "CSRF token mismatch":
                setCsrfToken();
                return "CSRF token mismatch"
            case `ERROR: duplicate key value violates unique constraint "uni_users_email" (SQLSTATE 23505)`:
                return "そちらのEmailはすでに登録されています";
            case "crypto/bcrypt: hashedPassword is not the hash of the given password":
                return "Passwordが間違っています";
            case "email: is not valid email format.":
                return "Emailの形式が間違っています";
            case "record not found":
                return "データが存在しません"
            default:
                return null;
        }
    }

    return { switchErrorHandling };
}