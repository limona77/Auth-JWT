import { createAsyncThunk } from "@reduxjs/toolkit";

import { AuthService } from "../../../services/auth";

import { IAuthResponse } from "../../../models/auth.ts";

import { AuthParams } from "./types.ts";

export const fetchLogin = createAsyncThunk<IAuthResponse, AuthParams>(
  "auth/fetchLogin",
  async (params) => {
    const { data } = await AuthService.login(params.email, params.password);
    localStorage.setItem("token", data.accessToken);
    return data;
  },
);
export const fetchRegister = createAsyncThunk<IAuthResponse, AuthParams>(
  "auth/fetchRegister",
  async (params) => {
    const { data } = await AuthService.register(params.email, params.password);
    localStorage.setItem("token", data.accessToken);
    return data;
  },
);
export const fetchAuthMe = createAsyncThunk("auth/fetchAuthMe", async () => {
  const token = localStorage.getItem("token") ?? "";
  const { data } = await AuthService.getUser(token);
  return data;
});
