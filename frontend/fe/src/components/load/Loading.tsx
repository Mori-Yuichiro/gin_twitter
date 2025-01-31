import { memo } from "react";

const Loading = memo(() => {
    return (
        <div className="min-h-screen">
            <svg className="animate-spin my-10 mx-auto" xmlns="http://www.w3.org/2000/svg" width="25" height="25" viewBox="0 0 24 24"><path fill="none" stroke="currentColor" strokeWidth="2" d="M13 4a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM7.34 6.34a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm11.32 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm0 11.32a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-11.32 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM21 12a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-8 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-8-8a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z" /></svg>
        </div>
    );
});

export default Loading;