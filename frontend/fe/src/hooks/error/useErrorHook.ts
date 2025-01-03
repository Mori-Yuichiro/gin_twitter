export const useErrorHook = () => {
    const switchErrorHandling = (msg: string) => {
        switch (msg) {
            case "missing csrf token in request header":
                // getCsrfToken();
                return "CSRFトークンがヘッダーに設定されていません";
            case "invalid csrf token":
                // getCsrfToken();
                return "CSRFトークンが間違っています";
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