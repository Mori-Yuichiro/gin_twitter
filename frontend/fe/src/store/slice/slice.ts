import { UserType } from "@/app/types/user";
import { createSlice, PayloadAction } from "@reduxjs/toolkit";

interface State {
    openSignUpModal: boolean;
    openLogInModal: boolean;
    currentUser: UserType | null;
    reload: boolean;
}

const initialState: State = {
    openSignUpModal: false,
    openLogInModal: false,
    currentUser: null,
    reload: false
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
        changeCurrentUser(state, action: PayloadAction<UserType>) {
            state.currentUser = action.payload;
        },
        toggleReload(state, action: PayloadAction<boolean>) {
            state.reload = action.payload;
        }
    }
});

export const {
    toggleSignupModal,
    toggleLogInModal,
    changeCurrentUser,
    toggleReload,
} = slice.actions;
export default slice.reducer;