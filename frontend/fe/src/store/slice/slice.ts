import { UserType } from "@/app/types/user";
import { createSlice, PayloadAction } from "@reduxjs/toolkit";

interface State {
    openSignUpModal: boolean;
    openLogInModal: boolean;
    openProfileModal: boolean;
    currentUser: UserType | null;
    reload: boolean;
    openTweetModal: boolean;
}

const initialState: State = {
    openSignUpModal: false,
    openLogInModal: false,
    openProfileModal: false,
    currentUser: null,
    reload: false,
    openTweetModal: false
};

const slice = createSlice({
    name: "state",
    initialState,
    reducers: {
        toggleSignupModal(state, action: PayloadAction<boolean>) {
            state.openSignUpModal = action.payload;
        },
        toggleLogInModal(state, action: PayloadAction<boolean>) {
            state.openLogInModal = action.payload;
        },
        toggleProfileModal(state, action: PayloadAction<boolean>) {
            state.openProfileModal = action.payload;
        },
        changeCurrentUser(state, action: PayloadAction<UserType>) {
            state.currentUser = action.payload;
        },
        toggleReload(state, action: PayloadAction<boolean>) {
            state.reload = action.payload;
        },
        toggleOpenTweetModal(state, action: PayloadAction<boolean>) {
            state.openTweetModal = action.payload;
        }
    }
});

export const {
    toggleSignupModal,
    toggleLogInModal,
    toggleProfileModal,
    changeCurrentUser,
    toggleReload,
    toggleOpenTweetModal
} = slice.actions;
export default slice.reducer;