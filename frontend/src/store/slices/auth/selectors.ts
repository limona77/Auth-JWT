import { RootState } from "../../index.ts";

export const selectIsAuth = (state: RootState) => state.auth;
