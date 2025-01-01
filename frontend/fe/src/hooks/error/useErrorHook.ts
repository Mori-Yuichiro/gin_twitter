export const useErrorHook = () => {
    const switchErrorHandling = (msg: string) => {
        switch (msg) {
            case 'ERROR: duplicate key value violates unique constraint \"uni_users_email\" (SQLSTATE 23505)':
                return "そちらのEmailはすでに登録されています";
        }
    }

    return { switchErrorHandling };
}