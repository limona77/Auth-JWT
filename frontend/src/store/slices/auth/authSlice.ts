import { createSlice, PayloadAction } from "@reduxjs/toolkit";

import { AuthState } from "./types.ts";
import { fetchLogin } from "./asyncActions.ts";

const initialState: AuthState = {
  isAuth: false,
};

export const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {
    setIsAuth(state, action: PayloadAction<boolean>) {
      state.isAuth = action.payload;
    },
  },
  extraReducers: (builder) => {
    builder.addCase(fetchLogin.fulfilled, (state) => {
      state.isAuth = true;
      alert("Вы успешно вошли в свой аккаунт!");
    });
    // builder.addCase(fetchLogin.rejected, () => {
    //   setIsAuth(false);
    // });
  },
});

export const { setIsAuth } = authSlice.actions;
export default authSlice.reducer;
