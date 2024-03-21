import axios from "axios";

export const Base_URL = "http://localhost:8000/";

export const httpInstance = axios.create({
  baseURL: Base_URL,
  withCredentials: true,
});

httpInstance.interceptors.response.use(
  (res) => {
    return res;
  },
  (err) => {
    if (axios.isAxiosError(err)) {
      alert(err.response?.data?.message);
    } else if (err instanceof Error) {
      alert(err.message);
    }
    return err;
  },
);
