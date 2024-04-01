import { createSlice } from "@reduxjs/toolkit";

import { AuthState } from "./types.ts";
import { fetchAuthMe, fetchLogin, fetchRegister } from "./asyncActions.ts";

const initialState: AuthState = {
  isAuth: false,
  isLoading: false,
};

export const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(fetchLogin.pending, (state) => {
      state.isLoading = true;
    });
    builder.addCase(fetchLogin.fulfilled, (state) => {
      state.isAuth = true;
      state.isLoading = false;
    });
    builder.addCase(fetchLogin.rejected, (state) => {
      state.isAuth = false;
      state.isLoading = false;
    });
    builder.addCase(fetchAuthMe.pending, (state) => {
      state.isLoading = true;
    });
    builder.addCase(fetchAuthMe.fulfilled, (state) => {
      state.isAuth = true;
      state.isLoading = false;
    });
    builder.addCase(fetchAuthMe.rejected, (state) => {
      state.isAuth = false;
      state.isLoading = false;
    });
    builder.addCase(fetchRegister.pending, (state) => {
      state.isLoading = true;
    });
    builder.addCase(fetchRegister.fulfilled, (state) => {
      state.isAuth = true;
      state.isLoading = false;
    });
    builder.addCase(fetchRegister.rejected, (state) => {
      state.isAuth = false;
      state.isLoading = false;
    });
  },
});

// export const {} = authSlice.actions;
export default authSlice.reducer;
