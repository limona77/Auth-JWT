import axios from "axios";

import { AuthService } from "../services/auth";

export const Base_URL = "http://localhost:8000/";

export const httpInstance = axios.create({
  baseURL: Base_URL,
  withCredentials: true,
});

httpInstance.interceptors.request.use((config) => {
  config.headers.Authorization = `Bearer ${localStorage.getItem("token")}`;
  return config;
});
let refresh = false;
httpInstance.interceptors.response.use(
  (res) => {
    return res;
  },
  async (err) => {
    const originalRequest = err.config;

    if (axios.isAxiosError(err)) {
      if (err.response?.status === 401 && err.config && !refresh) {
        refresh = true;
        const { data } = await AuthService.refresh();
        localStorage.setItem("token", data.accessToken);
        return httpInstance.request(originalRequest);
      }
      throw err.response?.data?.message;
    } else if (err instanceof Error) {
      throw err.message;
    }

    throw err;
  },
);
