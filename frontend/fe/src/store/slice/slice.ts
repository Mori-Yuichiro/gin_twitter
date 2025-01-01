import { createSlice, PayloadAction } from "@reduxjs/toolkit";

interface State {
    openSignUpModal: boolean;
}

const initialState: State = {
    openSignUpModal: false,
};

const slice = createSlice({
    name: "state",
    initialState,
    reducers: {
        toggleSignupModal(state, action: PayloadAction<boolean>) {
            state.openSignUpModal = action.payload;
        },
    }
});

export const {
    toggleSignupModal
} = slice.actions;
export default slice.reducer;