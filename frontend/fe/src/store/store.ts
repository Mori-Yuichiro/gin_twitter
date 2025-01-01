import { configureStore } from "@reduxjs/toolkit";
import sliceReducer from "@/store/slice/slice";

export const store = configureStore({
    reducer: {
        slice: sliceReducer
    }
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;