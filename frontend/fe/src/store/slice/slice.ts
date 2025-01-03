import { createSlice, PayloadAction } from "@reduxjs/toolkit";

interface State {
    openSignUpModal: boolean;
    openLogInModal: boolean;
}

const initialState: State = {
    openSignUpModal: false,
    openLogInModal: false
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
    }
});

export const {
    toggleSignupModal,
    toggleLogInModal
} = slice.actions;
export default slice.reducer;