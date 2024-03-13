import axios from "axios";

export const Base_URL = "http://localhost:8000/";

export const axiosInstance = axios.create({
  baseURL: Base_URL,
  withCredentials: true,
});
axios.interceptors.response.use(
  function (response) {
    return response;
  },
  function (error) {
    return Promise.reject(error);
  },
);
